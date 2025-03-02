package user

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	"clean-hex/pkg/framwork/helpers/is"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"gorm.io/gorm"
)

type CreateUserCommandHandler struct {
	uow internal.UnitOfWorkImp
}

func NewCreateUserCommandHandler(uow internal.UnitOfWorkImp) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{uow: uow}
}

func (u *CreateUserCommandHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(*domain.CreateUserCommand)

	return u.uow.Do(ctx, func(ctx context.Context, tx *gorm.DB) (any, error) {
		_, err := u.uow.User().FindByUserName(ctx, command.UserName)
		if !is.Error(err, gorm.ErrRecordNotFound) {
			return nil, errors.BadRequest("User.AlreadyExists")
		}

		user, err := entities.NewUser(command.UserName, command.Age, command.Amount)
		if !is.Empty(err) {
			return nil, err
		}

		err = u.uow.User().Save(ctx, user)
		if !is.Empty(err) {
			return nil, errors.BadRequest("Operation.CanNot")
		}

		return user, nil
	})

}
