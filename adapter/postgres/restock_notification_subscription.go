package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RestockNotificationSubscription struct {
	bun.BaseModel `bun:"table:restock_notification_subscriptions,alias:rns"`

	// Attributes.
	ID        uuid.UUID `bun:",pk"`
	UserID    uuid.UUID
	ProductID int64
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations.
	User    *User    `bun:"rel:belongs-to"`
	Product *Product `bun:"rel:belongs-to"`
}
