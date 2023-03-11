package subscription

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/alextanhongpin/restocknotif/adapter/postgres"
	"github.com/alextanhongpin/restocknotif/types/slices"
	"github.com/google/uuid"
)

type Repository struct {
	db postgres.IDB
}

func NewRepository(db postgres.IDB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(ctx context.Context, userID uuid.UUID) ([]Subscription, error) {
	u := postgres.User{
		ID: userID,
	}
	err := r.db.NewSelect().
		Model(&u).
		Relation("Subscriptions").
		WherePK().
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to list restock notification subscriptions: %w", err)
	}

	return slices.Map(u.Subscriptions, func(i int) Subscription {
		s := u.Subscriptions[i]
		return Subscription{
			ID:        s.ID,
			UserID:    s.UserID,
			ProductID: s.ProductID,
			Quantity:  s.Quantity,
		}
	}), nil
}

func (r *Repository) Create(ctx context.Context, userID uuid.UUID, productID, quantity int64) error {
	m := postgres.RestockNotificationSubscription{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	_, err := r.db.NewInsert().
		Model(&m).
		Column(
			"user_id",
			"product_id",
			"quantity",
		).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create restock notification subscription: %w", err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, userID, restockNotificationSubscriptionID uuid.UUID, quantity int64) error {
	m := postgres.RestockNotificationSubscription{
		ID:       restockNotificationSubscriptionID,
		UserID:   userID,
		Quantity: quantity,
	}
	_, err := r.db.NewUpdate().
		Model(&m).
		Column("quantity").
		WherePK().
		Where("user_id = ?user_id").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update restock notification subscription: %w", err)
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, userID, restockNotificationSubscriptionID uuid.UUID) error {
	m := postgres.RestockNotificationSubscription{
		ID:     restockNotificationSubscriptionID,
		UserID: userID,
	}
	_, err := r.db.NewDelete().
		Model(&m).
		WherePK().
		Where("user_id = ?user_id").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete restock notification subscription: %w", err)
	}

	return nil
}
