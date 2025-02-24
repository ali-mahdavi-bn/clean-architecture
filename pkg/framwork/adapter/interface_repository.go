package adapter

import (
	"context"
)

type BaseRepository[E Entity] interface {
	FindByID(ctx context.Context, id uint) (E, error)
	FindByField(ctx context.Context, field string, value interface{}) (E, error)
	Remove(ctx context.Context, model E) error
	Save(ctx context.Context, model E) error
}
