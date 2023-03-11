package api

import (
	"net/http"

	"github.com/alextanhongpin/restocknotif/rest/middleware"
	"github.com/alextanhongpin/restocknotif/rest/response"
	"github.com/alextanhongpin/restocknotif/types/slices"
	"github.com/alextanhongpin/restocknotif/usecase/subscription"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type SubscriptionAPI struct {
	useCase subscription.T
}

func NewSubscriptionAPI(uc subscription.T) *SubscriptionAPI {
	return &SubscriptionAPI{
		useCase: uc,
	}
}

func (api *SubscriptionAPI) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := middleware.MustUserIDFromContext(ctx)
	subscriptions, err := api.useCase.List(ctx, userID)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	data := slices.Map(subscriptions, func(i int) Subscription {
		s := subscriptions[i]
		return *NewSubscription(s)
	})

	response.Success(w, data, http.StatusOK)
}

func (api *SubscriptionAPI) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := middleware.MustUserIDFromContext(ctx)
	req, err := response.ParseBody[PostCreateRequest](r)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	err = api.useCase.Create(ctx, userID, req.ProductID, req.Quantity)
	if err != nil {
		response.Failure(w, err, http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *SubscriptionAPI) Patch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := middleware.MustUserIDFromContext(ctx)
	req, err := response.ParseBody[PatchRequest](r)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	subscriptionIDStr := chi.URLParam(r, "subscription_id")
	subscriptionID, err := uuid.Parse(subscriptionIDStr)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	err = api.useCase.Update(ctx, userID, subscriptionID, req.Quantity)
	if err != nil {
		response.Failure(w, err, http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *SubscriptionAPI) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := middleware.MustUserIDFromContext(ctx)
	subscriptionIDStr := chi.URLParam(r, "subscription_id")
	subscriptionID, err := uuid.Parse(subscriptionIDStr)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	err = api.useCase.Delete(ctx, userID, subscriptionID)
	if err != nil {
		response.Failure(w, err, http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
