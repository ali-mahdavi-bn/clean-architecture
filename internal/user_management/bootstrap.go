package user_management

import (
	"clean-hex/internal/user_management/adapter/repositories"
	"clean-hex/internal/user_management/domain"
	handlers "clean-hex/internal/user_management/service_layer/handlers/user"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"gorm.io/gorm"
)

func Bootstrap(db *gorm.DB) *messagebus.MessageBus {
	user_repo := repositories.NewUserGormRepository(db)

	bus := messagebus.NewMessageBus(db)
	bus.Register(domain.CreateUserCommand{}, handlers.NewCreateUserHandler(user_repo))
	bus.Register(domain.UpdateUserCommand{}, handlers.NewUpdateUserHandler(user_repo))
	bus.Register(domain.DeleteUserCommand{}, handlers.NewDeleteUserHandler(user_repo))

	return bus
}
