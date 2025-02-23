package adapter

import (
	"context"
)

type BaseRepository[E Entity] interface {
	ByID(ctx context.Context, id uint) (E, error)
	ByField(ctx context.Context, field string, value interface{}) (E, error)
	Add(ctx context.Context, model E) error
}
