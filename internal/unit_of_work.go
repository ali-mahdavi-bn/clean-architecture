package internal

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/pkg/framwork/service_layer/types"
	"clean-hex/pkg/framwork/service_layer/unit_of_work"
	"context"
	"gorm.io/gorm"
)

var (
	_ UnitOfWorkImp = &unitOfWorkImp{}
)

type UnitOfWorkImp interface {
	unit_of_work.UnitOfWork
	Register()
	User() repositories.UserRepository
	Trade() repositories.TradeRepository
}

type unitOfWorkImp struct {
	unit_of_work.UnitOfWork
	user  repositories.UserRepository
	trade repositories.TradeRepository
}

func NewGormUnitOfWorkImp(db *gorm.DB) UnitOfWorkImp {
	return &unitOfWorkImp{UnitOfWork: unit_of_work.NewGormUnitOfWork(db)}
}
func (uow *unitOfWorkImp) Register() {
	uow.user = repositories.NewUserGormRepository(uow.UnitOfWork.GetSession())
	uow.trade = repositories.NewTradeGormRepository(uow.UnitOfWork.GetSession())
}
func (uow *unitOfWorkImp) Do(ctx context.Context, fc types.UowUseCase) (result interface{}, err error) {
	defer func() {
		if recover() != nil || err != nil {
			_ = uow.Rollback()
		}
	}()

	if uow.Begin() != nil {
		return nil, err
	}
	// initial repository
	uow.Register()

	result, err = uow.UnitOfWork.Do(ctx, fc)
	if err != nil {
		return nil, err
	}
	if err = uow.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (uow *unitOfWorkImp) User() repositories.UserRepository {
	return uow.user
}

func (uow *unitOfWorkImp) Trade() repositories.TradeRepository {
	return uow.trade
}
