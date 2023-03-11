package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/alextanhongpin/restocknotif/rest/response"
	"github.com/alextanhongpin/restocknotif/usecase/authn"
	"github.com/google/uuid"
)

type authenticator interface {
	CreateToken(userID uuid.UUID) (string, error)
}

type AuthAPI struct {
	useCase       authn.T
	authenticator authenticator
}

func NewAuthAPI(uc authn.T, a authenticator) *AuthAPI {
	return &AuthAPI{
		useCase:       uc,
		authenticator: a,
	}
}

func (api *AuthAPI) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := response.ParseBody[LoginRequest](r)
	if err != nil {
		response.Failure(w, err, http.StatusBadRequest)
		return
	}

	user, err := api.useCase.Find(ctx, req.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Failure(w, err, http.StatusNotFound)
		} else {
			response.Failure(w, err, http.StatusConflict)
		}

		return
	}

	accessToken, err := api.authenticator.CreateToken(user.ID)
	if err != nil {
		response.Failure(w, err, http.StatusPreconditionFailed)
		return
	}

	data := LoginResponse{
		AccessToken: accessToken,
	}

	response.Success(w, data, http.StatusOK)
}
