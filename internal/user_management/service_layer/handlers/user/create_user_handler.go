package handlers

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/errors"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
)

type CreateUserHandler struct {
	userRepo repositories.UserRepository
}

func NewCreateUserHandler(userRepo repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{userRepo: userRepo}
}

func (u *CreateUserHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(domain.CreateUserCommand)

	_, err := u.userRepo.FindByUserName(ctx, command.UserName)
	if err == nil {
		return nil, errors.BadRequest("User.AlreadyExists")
	}

	user, err := entities.NewUser(command.UserName, command.Age)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Save(ctx, user)
	if err != nil {
		return nil, errors.BadRequest("CanNot.Operation")
	}

	return user, nil
}
