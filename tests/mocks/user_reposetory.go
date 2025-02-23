package mocks

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain/entities"
	"errors"

	"context"
)

type FakeUserRepository struct {
	FakRepository[*entities.User]
}

func (f *FakeUserRepository) ByUserName(ctx context.Context, username string) (*entities.User, error) {
	args := f.Called(ctx, username)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func NewFakeUserRepository() repositories.UserRepository {
	ctx := context.Background()
	userRepo := &FakeUserRepository{
		FakRepository: *NewFakeRepository[*entities.User](),
	}
	userRepo.On("Add", ctx, &entities.User{UserName: "ali", Age: 20}).Return(nil)
	userRepo.On("Add", ctx, &entities.User{UserName: "NewAli", Age: 20}).Return(nil)
	userRepo.On("ByUserName", ctx, "ali").Return(&entities.User{UserName: "ali", Age: 20}, nil)
	userRepo.On("ByUserName", ctx, "NewAli").Return((*entities.User)(nil), errors.New("User.NotFound"))
	userRepo.On("ByUserName", ctx, "Bob").Return((*entities.User)(nil), errors.New("User.AlreadyExists"))
	return userRepo
}
