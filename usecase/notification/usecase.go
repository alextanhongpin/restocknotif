package notification

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type T interface {
	SendRestockNotification(ctx context.Context, userID, subscriptionID int64) error
}

type UseCase struct {
	weaver.Implements[T]
}

func (uc *UseCase) Init(ctx context.Context) error {
	// Register subscriber.
	return nil
}

func (uc *UseCase) BatchSendRestockNotification(ctx context.Context, itemID, qty int64) error {
	// query users
	// for each user, send the restock notification.
	// update the status in db.
	return nil
}

func (uc *UseCase) SendRestockNotification(ctx context.Context, userID, subscriptionID int64) error {
	// uc.send
	return nil
}
