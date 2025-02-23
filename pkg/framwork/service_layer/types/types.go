package types

import (
	"context"
)

type Command interface{}

type HandlerType interface {
	Handle(ctx context.Context, cmd Command) (any, error)
}

type Modules interface {
	Init() error
}
