package postgres

import (
	"time"

	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	// Attributes.
	ID        int64 `bun:",pk"`
	Name      string
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations.
	Subscriptions []RestockNotificationSubscription `bun:"rel:has-many,join:id=product_id"`
}
