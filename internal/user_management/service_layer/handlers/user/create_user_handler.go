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
	user_repo repositories.UserRepository
}

func NewCreateUserHandler(user_repo repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{user_repo: user_repo}
}

func (u *CreateUserHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(domain.CreateUserCommand)

	_, err := u.user_repo.ByUserName(ctx, command.UserName)
	if err == nil {
		return nil, errors.BadRequest("User.AlreadyExists")
	}

	user, err := entities.NewUser(command.UserName, command.Age)
	if err != nil {
		return nil, err
	}

	err = u.user_repo.Add(ctx, user)
	if err != nil {
		return nil, errors.BadRequest("CanNot.Operation")
	}

	return user, nil
}
