package user_management

import (
	unit_of_work "clean-hex/internal"
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/service_layer/handlers/trade"
	"clean-hex/internal/user_management/service_layer/handlers/user"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"gorm.io/gorm"
)

func Bootstrap(db *gorm.DB) *messagebus.MessageBus {
	uow := unit_of_work.NewGormUnitOfWorkImp(db)
	bus := messagebus.NewMessageBus(uow)

	// user
	bus.Register(domain.CreateUserCommand{}, user.NewCreateUserCommandHandler(bus.Uow))
	bus.Register(domain.UpdateUserCommand{}, user.NewUpdateUserHandler(bus.Uow))
	bus.Register(domain.DeleteUserCommand{}, user.NewDeleteUserHandler(bus.Uow))

	// trade
	bus.Register(domain.CreateTradeCommand{}, trade.NewCreateTradeCommandHandler(bus.Uow))
	bus.Register(domain.UpdateTradeCommand{}, trade.NewUpdateTradeHandler(bus.Uow))
	bus.Register(domain.DeleteTradeCommand{}, trade.NewDeleteTradeHandler(bus.Uow))

	return bus
}
