package messagebus

import (
	"clean-hex/internal"
	"clean-hex/pkg/framwork/helpers/is"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"errors"
	"reflect"
)

var (
	handlerNotFountError = errors.New("no handler found for command")
	handlerInvalidError  = errors.New("invalid handler type")
)

type MessageBus struct {
	Uow      internal.UnitOfWorkImp
	handlers map[string]types.HandlerType
}

func NewMessageBus(uow internal.UnitOfWorkImp) *MessageBus {
	return &MessageBus{
		handlers: make(map[string]types.HandlerType),
		Uow:      uow,
	}
}

func (m *MessageBus) Register(cmd types.Command, handler types.HandlerType) {
	m.handlers[reflect.TypeOf(cmd).String()] = handler
}

func (m *MessageBus) Handle(ctx context.Context, cmd types.Command) (any, error) {
	typeCmd := reflect.TypeOf(cmd)
	if is.Ptr(cmd) {
		typeCmd = typeCmd.Elem()
	}
	typeName := typeCmd.String()
	handler, exists := m.handlers[typeName]
	if !exists {
		return nil, handlerNotFountError
	} else if h, ok := handler.(types.HandlerType); ok {
		return h.Handle(ctx, cmd)
	}

	return nil, handlerInvalidError
}
