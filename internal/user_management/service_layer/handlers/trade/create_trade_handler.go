package trade

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	"clean-hex/pkg/framwork/helpers/is"
	kafka "clean-hex/pkg/framwork/infrastructure/kafak"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"gorm.io/gorm"
)

type CreateTradeCommandHandler struct {
	uow internal.UnitOfWorkImp
}

func NewCreateTradeCommandHandler(uow internal.UnitOfWorkImp) *CreateTradeCommandHandler {
	return &CreateTradeCommandHandler{uow: uow}
}

func (u *CreateTradeCommandHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(*domain.CreateTradeCommand)

	return u.uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {
		trade, err := entities.NewTrade(command.UserId, command.Stock, command.Price, command.Amount)
		if !is.Empty(err) {
			return nil, err
		}
		err = u.uow.Trade().Save(ctx, trade)
		if !is.Empty(err) {
			return nil, errors.BadRequest("Operation.CanNot")
		}

		if err != nil {
			return nil, errors.BadRequest("Operation.CanNot")
		}
		if !is.Empty(kafka.Service.SendMessage(kafka.KAFKA_TOPIC_UPDATE_TRADE, trade)) {
			return nil, errors.BadRequest("Operation.CanNot")
		}

		return trade, nil
	})

}
