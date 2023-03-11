package authn

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/alextanhongpin/restocknotif/adapter/postgres"
	"github.com/google/uuid"
)

type User struct {
	weaver.AutoMarshal

	ID   uuid.UUID
	Name string
}

func NewUser(u postgres.User) *User {
	return &User{
		ID:   u.ID,
		Name: u.Name,
	}
}
