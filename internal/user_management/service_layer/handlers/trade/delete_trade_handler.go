package trade

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/domain"
	"clean-hex/pkg/framwork/errors"
	"clean-hex/pkg/framwork/helpers/is"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"gorm.io/gorm"
)

type DeleteTradeHandler struct {
	uow internal.UnitOfWorkImp
}

func NewDeleteTradeHandler(uow internal.UnitOfWorkImp) *DeleteTradeHandler {
	return &DeleteTradeHandler{uow: uow}
}

func (u *DeleteTradeHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(*domain.DeleteUserCommand)

	return u.uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (interface{}, error) {
		user, err := u.uow.User().FindByID(ctx, command.UserId)
		if !is.Error(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("User.NotFound")
		}

		if err = u.uow.User().Remove(ctx, user); !is.Empty(err) {
			return nil, errors.BadRequest("Operation.CanNot")
		}
		return nil, nil
	})
}
