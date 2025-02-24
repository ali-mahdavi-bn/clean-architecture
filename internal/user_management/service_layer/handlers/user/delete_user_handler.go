package handlers

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain"
	"clean-hex/pkg/errors"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
)

type DeleteUserHandler struct {
	userRepo repositories.UserRepository
}

func NewDeleteUserHandler(userRepo repositories.UserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{userRepo: userRepo}
}

func (u *DeleteUserHandler) Handle(ctx context.Context, cmd types.Command) (any, error) {
	command := cmd.(domain.DeleteUserCommand)

	user, err := u.userRepo.FindByID(ctx, command.UserId)
	if err != nil && user.IsDeleted() {
		return nil, errors.NotFound("User.NotFound")
	}
	if err = u.userRepo.Remove(ctx, user); err != nil {
		return nil, errors.BadRequest("CanNot.Operation")
	}
	return nil, nil

}
