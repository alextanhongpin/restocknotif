package subscription

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/adapter/postgres"
	"github.com/google/uuid"
)

type T interface {
	Create(ctx context.Context, userID uuid.UUID, productID, quantity int64) error
	Update(ctx context.Context, userID, subscriptionID uuid.UUID, quantity int64) error
	Delete(ctx context.Context, userID, subscriptionID uuid.UUID) error
	List(ctx context.Context, userID uuid.UUID) ([]Subscription, error)
}

type repository interface {
	List(ctx context.Context, userID uuid.UUID) ([]Subscription, error)
	Create(ctx context.Context, userID uuid.UUID, productID, quantity int64) error
	Update(ctx context.Context, userID, restockNotificationSubscriptionID uuid.UUID, quantity int64) error
	Delete(ctx context.Context, userID, restockNotificationSubscriptionID uuid.UUID) error
}

type UseCase struct {
	weaver.Implements[T]
	repo repository
}

func (uc *UseCase) Init(ctx context.Context) error {
	uc.repo = NewRepository(postgres.New())

	return nil
}

func (uc *UseCase) List(ctx context.Context, userID uuid.UUID) ([]Subscription, error) {
	subscriptions, err := uc.repo.List(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list subscriptions: %w", err)
	}

	return subscriptions, nil
}

func (uc *UseCase) Create(ctx context.Context, userID uuid.UUID, productID, quantity int64) error {
	err := uc.repo.Create(ctx, userID, productID, quantity)
	if err != nil {
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}

func (uc *UseCase) Update(ctx context.Context, userID, subscriptionID uuid.UUID, quantity int64) error {
	err := uc.repo.Update(ctx, userID, subscriptionID, quantity)
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	return nil
}

func (uc *UseCase) Delete(ctx context.Context, userID, subscriptionID uuid.UUID) error {
	err := uc.repo.Delete(ctx, userID, subscriptionID)
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}
