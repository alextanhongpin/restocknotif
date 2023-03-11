package authn

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

func (r *Repository) FindByName(ctx context.Context, name string) (*User, error) {
	u := postgres.User{
		Name: name,
	}

	err := r.db.NewSelect().
		Model(&u).
		Where("name = ?name").
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return NewUser(u), nil
}
