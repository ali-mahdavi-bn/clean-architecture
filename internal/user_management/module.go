package user_management

import (
	"clean-hex/internal/user_management/adapter/model"
	"clean-hex/internal/user_management/entryporint"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserManagementModule struct {
	Ctx         context.Context
	DB          *gorm.DB
	RouterGroup *gin.RouterGroup
}

func (u *UserManagementModule) AutoMigration() error {
	return u.DB.AutoMigrate(new(model.User))
}

func (u *UserManagementModule) Init() error {
	err := u.AutoMigration()
	if err != nil {
		return err
	}
	bus := Bootstrap(u.DB)
	entryporint.RegisterV1Routers(bus, u.RouterGroup)

	return nil
}
