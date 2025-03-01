package mocks

import (
	"clean-hex/internal"
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var _ internal.UnitOfWorkImp = &FakeUnitOfWork{}

type FakeUnitOfWork struct {
	mock.Mock
	tx    *gorm.DB
	user  repositories.UserRepository
	trade repositories.TradeRepository
}

func (f *FakeUnitOfWork) Trade() repositories.TradeRepository {
	return f.trade
}

func (f *FakeUnitOfWork) Register() {
	f.user = NewFakeUserRepository()
}

func (f *FakeUnitOfWork) User() repositories.UserRepository {
	return f.user
}

func (f *FakeUnitOfWork) GetSession() *gorm.DB {
	args := f.Called()
	return args.Get(0).(*gorm.DB)

}

func NewFakeUnitOfWork() *FakeUnitOfWork {
	fakeUnitOfWork := &FakeUnitOfWork{}
	fakeUnitOfWork.On("Begin").Return(nil)
	fakeUnitOfWork.On("Commit").Return(nil)
	fakeUnitOfWork.On("Rollback").Return(nil)

	return fakeUnitOfWork
}

func (f *FakeUnitOfWork) Begin() error {
	args := f.Called()
	return args.Error(0)
}

func (f *FakeUnitOfWork) Do(ctx context.Context, fc types.UowUseCase) (result interface{}, err error) {
	defer func() {
		if a := recover(); a != nil || err != nil {
			_ = f.Rollback()
		}
	}()

	if f.Begin() != nil {
		return nil, err
	}
	f.Register()
	result, err = fc(ctx, f.tx)
	if err != nil {

		return nil, err
	}

	if err = f.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (f *FakeUnitOfWork) Commit() error {
	args := f.Called()
	return args.Error(0)
}

func (f *FakeUnitOfWork) Rollback() error {
	args := f.Called()
	return args.Error(0)
}
