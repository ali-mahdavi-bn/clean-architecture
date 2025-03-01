// uow.go
package unit_of_work

import (
	"clean-hex/pkg/framwork/service_layer/types"
	"context"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Begin() error
	Do(ctx context.Context, fc types.UowUseCase) (result interface{}, err error)
	Commit() error
	Rollback() error
	GetSession() *gorm.DB
}

type GormUnitOfWork struct {
	DB *gorm.DB
	tx *gorm.DB
}

func NewGormUnitOfWork(db *gorm.DB) UnitOfWork {
	return &GormUnitOfWork{
		DB: db,
	}
}

func (uow *GormUnitOfWork) GetSession() *gorm.DB {
	return uow.tx
}

func (uow *GormUnitOfWork) Begin() error {
	uow.tx = uow.DB.Begin()
	if uow.tx.Error != nil {
		return uow.tx.Error
	}
	return nil
}

func (uow *GormUnitOfWork) Do(ctx context.Context, fc types.UowUseCase) (result interface{}, err error) {
	return fc(ctx, uow.tx)
}

func (uow *GormUnitOfWork) Commit() error {
	return uow.tx.Commit().Error
}

func (uow *GormUnitOfWork) Rollback() error {
	return uow.tx.Rollback().Error
}
