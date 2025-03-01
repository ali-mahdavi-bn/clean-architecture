package mocks

import (
	"clean-hex/internal/user_management"
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/service_layer/handlers/user"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"gorm.io/gorm"
)

func MockUserManagementBootstrapTestApp() *FakeMessageBus {
	bus := NewFakeMessageBus(NewFakeUnitOfWork())
	bus.Register(domain.CreateUserCommand{}, user.NewCreateUserCommandHandler(bus.Uow))
	return bus
}

func SqliteUserManagementBootstrapTestApp(db *gorm.DB) *messagebus.MessageBus {
	return user_management.Bootstrap(db)
}
