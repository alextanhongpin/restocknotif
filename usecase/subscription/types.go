package subscription

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Subscription struct {
	weaver.AutoMarshal

	ID        uuid.UUID
	UserID    uuid.UUID
	ProductID int64
	Quantity  int64
}
