package handlers

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain"
	"clean-hex/pkg/errors"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
)

type UpdateUserHandler struct {
	userRepo repositories.UserRepository
}

func NewUpdateUserHandler(user_repo repositories.UserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{userRepo: user_repo}
}

func (u *UpdateUserHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(domain.UpdateUserCommand)
	user, err := u.userRepo.FindByID(ctx, command.UserId)
	if err != nil || user.IsDeleted() {
		return nil, errors.NotFound("User.NotFound")
	}
	if _, err = u.userRepo.FindByUsernameExcludingID(ctx, command.UserName, command.UserId); err == nil {
		return nil, errors.BadRequest("User.AlreadyExists")
	}

	if err = user.Update(command.UserName, command.Age); err != nil {
		return nil, err
	}
	if err = u.userRepo.Save(ctx, user); err != nil {
		return nil, errors.BadRequest("CanNot.Operation")
	}
	return user, nil
}
