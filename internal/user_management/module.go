package user_management

import (
	unit_of_work "clean-hex/internal"
	"clean-hex/internal/user_management/adapter/model"
	"clean-hex/internal/user_management/entryporint"
	"clean-hex/internal/user_management/workers"
	"clean-hex/pkg/framwork/service_layer/cache"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserManagementModule struct {
	Ctx         context.Context
	DB          *gorm.DB
	RedisStore  cache.Store
	RouterGroup *gin.RouterGroup
	uow         unit_of_work.UnitOfWorkImp
}

func (u *UserManagementModule) AutoMigration() error {
	return u.DB.AutoMigrate(
		new(model.User),
		new(model.Trade),
	)
}
func (u *UserManagementModule) StartWorker() {

	go workers.CacheTrade(u.uow, u.RedisStore)

}

func (u *UserManagementModule) Init() error {
	err := u.AutoMigration()
	if err != nil {
		return err
	}
	bus := Bootstrap(u.DB)
	u.uow = bus.Uow
	entryporint.RegisterV1Routers(bus, u.RouterGroup, u.RedisStore)

	u.StartWorker()

	return nil
}
