package messagebus

import (
	"clean-hex/pkg/framwork/service_layer/types"
	"clean-hex/pkg/helpers"
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
)

var (
	handlerNotFountError = errors.New("no handler found for command")
	handlerInvalidError  = errors.New("invalid handler type")
)

type MessageBus struct {
	DB       *gorm.DB
	handlers map[string]types.HandlerType
}

func NewMessageBus(db *gorm.DB) *MessageBus {
	return &MessageBus{
		handlers: make(map[string]types.HandlerType),
		DB:       db,
	}
}

func (m *MessageBus) Register(cmd types.Command, handler types.HandlerType) {
	typeName := reflect.TypeOf(cmd).String()
	m.handlers[typeName] = handler
}

func (m *MessageBus) Handle(ctx context.Context, cmd types.Command) (any, error) {
	typeCmd := reflect.TypeOf(cmd)
	if helpers.IsPtr(typeCmd) {
		typeCmd = typeCmd.Elem()
	}
	typeName := typeCmd.String()
	handler, exists := m.handlers[typeName]
	if !exists {
		return nil, handlerNotFountError
	}
	if h, ok := handler.(types.HandlerType); ok {
		return h.Handle(ctx, cmd)
	}

	return nil, handlerInvalidError
}
