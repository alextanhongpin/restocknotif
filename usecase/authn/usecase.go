package authn

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/adapter/postgres"
)

type T interface {
	Find(ctx context.Context, name string) (*User, error)
}

type repository interface {
	FindByName(ctx context.Context, name string) (*User, error)
}

type UseCase struct {
	weaver.Implements[T]

	repo repository
}

func (uc *UseCase) Init(ctx context.Context) error {
	uc.repo = NewRepository(postgres.New())

	return nil
}

func (uc *UseCase) Find(ctx context.Context, name string) (*User, error) {
	user, err := uc.repo.FindByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by name %q: %w", name, err)
	}

	return user, nil
}
