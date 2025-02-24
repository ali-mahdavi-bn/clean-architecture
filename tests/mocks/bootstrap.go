package mocks

import (
	"clean-hex/internal/user_management"
	"clean-hex/internal/user_management/domain"
	handlers "clean-hex/internal/user_management/service_layer/handlers/user"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"gorm.io/gorm"
)

func MockUserManagementBootstrapTestApp() *messagebus.MessageBus {
	user_repo := NewFakeUserRepository()

	bus := messagebus.NewMessageBus(nil)
	bus.Register(domain.CreateUserCommand{}, handlers.NewCreateUserHandler(user_repo))

	return bus
}

func SqliteUserManagementBootstrapTestApp(db *gorm.DB) *messagebus.MessageBus {
	return user_management.Bootstrap(db)
}
