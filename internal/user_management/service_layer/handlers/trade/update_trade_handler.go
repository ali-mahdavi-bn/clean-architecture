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

type UpdateTradeHandler struct {
	uow internal.UnitOfWorkImp
}

func NewUpdateTradeHandler(uow internal.UnitOfWorkImp) *UpdateTradeHandler {
	return &UpdateTradeHandler{uow: uow}
}

func (u *UpdateTradeHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(*domain.UpdateUserCommand)
	return u.uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {

		user, err := u.uow.User().FindByID(ctx, command.UserId)
		if is.Error(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("User.NotFound")
		}

		_, err = u.uow.User().FindByUsernameExcludingID(ctx, command.UserName, command.UserId)
		if !is.Error(err, gorm.ErrRecordNotFound) {
			return nil, errors.BadRequest("User.AlreadyExists")
		}

		if err = user.Update(command.UserName, command.Age, command.Amount); !is.Empty(err) {
			return nil, err
		}
		if !is.Empty(u.uow.User().Save(ctx, user)) {
			return nil, errors.BadRequest("Operation.CanNot")
		}
		return user, nil
	})
}
