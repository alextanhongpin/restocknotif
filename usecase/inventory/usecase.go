package product

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/adapter/postgres"
)

type T interface {
	IncrementStock(ctx context.Context, productID, quantity int64) error
}

type repository interface {
	IncrementStock(ctx context.Context, productID, quantity int64) error
}

type UseCase struct {
	weaver.Implements[T]
	repo repository
}

func (uc *UseCase) Init(ctx context.Context) error {
	uc.Logger().Debug("Init")
	uc.repo = NewRepository(postgres.New())

	return nil
}

// IncrementStock increments the product stock by the given quantity.
func (uc *UseCase) IncrementStock(ctx context.Context, productID, quantity int64) error {
	uc.Logger().Debug("IncrementStock",
		"productID", productID,
		"quantity", quantity,
	)

	err := uc.repo.IncrementStock(ctx, productID, quantity)
	if err != nil {
		return fmt.Errorf("failed to increment stock: %w", err)
	}

	return nil
}
