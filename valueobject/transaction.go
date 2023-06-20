package valueobject

import (
	"time"

	"github.com/google/uuid"
)

//immutable

type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
