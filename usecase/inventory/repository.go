package product

import (
	"context"

	"github.com/alextanhongpin/restocknotif/adapter/postgres"
)

type Repository struct {
	db postgres.IDB
}

func NewRepository(db postgres.IDB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) IncrementStock(ctx context.Context, productID, quantity int64) error {
	p := postgres.Product{
		Quantity: quantity,
		ID:       productID,
	}

	_, err := r.db.NewUpdate().
		Model(&p).
		Column("quantity").
		Set("quantity = quantity + ?quantity").
		WherePK().
		Exec(ctx)

	return err
}
