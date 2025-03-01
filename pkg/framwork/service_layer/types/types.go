package types

import (
	"context"
	"gorm.io/gorm"
)

type Command interface{}

type HandlerType interface {
	Handle(ctx context.Context, cmd Command) (any, error)
}
type RedisUseCase func(ctx context.Context) (interface{}, error)
type UowUseCase func(ctx context.Context, tx *gorm.DB) (interface{}, error)

type Modules interface {
	Init() error
}
