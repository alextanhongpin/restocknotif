package api

import (
	"github.com/alextanhongpin/restocknotif/usecase/subscription"
	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	ProductID int64     `json:"productId"`
	Quantity  int64     `json:"quantity"`
}

func NewSubscription(e subscription.Subscription) *Subscription {
	return &Subscription{
		ID:        e.ID,
		UserID:    e.UserID,
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}

type PostCreateRequest struct {
	ProductID int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

type PatchRequest struct {
	Quantity int64 `json:"quantity"`
}
