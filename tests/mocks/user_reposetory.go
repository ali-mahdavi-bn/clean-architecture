package mocks

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain/entities"
	"errors"
	"gorm.io/gorm"

	"context"
)

type FakeUserRepository struct {
	FakRepository[*entities.User]
}

func (f *FakeUserRepository) FindByUserName(ctx context.Context, username string) (*entities.User, error) {
	args := f.Called(ctx, username)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.User), args.Error(1)
	}
	return nil, args.Error(1)
}
func (f *FakeUserRepository) FindByUsernameExcludingID(ctx context.Context, username string, id uint) (*entities.User, error) {
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
	userRepo.On("Save", ctx, &entities.User{UserName: "ali", Age: 20}).Return(nil)
	userRepo.On("Save", ctx, &entities.User{UserName: "NewAli", Age: 20}).Return(nil)
	userRepo.On("FindByUserName", ctx, "ali").Return(&entities.User{UserName: "ali", Age: 20}, nil)
	userRepo.On("FindByUserName", ctx, "NewAli").Return((*entities.User)(nil), gorm.ErrRecordNotFound)
	userRepo.On("FindByUserName", ctx, "Bob").Return((*entities.User)(nil), gorm.ErrRecordNotFound)
	userRepo.On("FindByUsernameExcludingID", ctx, "Bob").Return((*entities.User)(nil), errors.New("User.AlreadyExists"))
	return userRepo
}
