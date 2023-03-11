package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// Attributes.
	ID        uuid.UUID `bun:",pk"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Subscriptions []RestockNotificationSubscription `bun:"rel:has-many,join:id=user_id"`
}
