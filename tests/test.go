package tests

import (
	"clean-hex/internal/user_management/domain"
	handlers "clean-hex/internal/user_management/service_layer/handlers/user"
	"clean-hex/pkg/framwork/infrastructure/databases"
	"clean-hex/pkg/framwork/service_layer/messagebus"
	"clean-hex/tests/mocks"
)

func MockUserManagementBootstrapTestApp() *messagebus.MessageBus {
	user_repo := mocks.NewFakeUserRepository()

	bus := messagebus.NewMessageBus(nil)
	bus.Register(domain.CreateUserCommand{}, handlers.NewCreateUserHandler(user_repo))

	return bus
}

func SqliteUserManagementBootstrapTestApp() *messagebus.MessageBus {
	db, _ := databases.NewDbConnection(databases.Config{
		Debug:        true,
		DBType:       "sqlite3",
		DSN:          "./test.db",
		MaxLifetime:  1,
		MaxIdleTime:  1,
		MaxIdleConns: 1,
		MaxOpenConns: 1,
		TablePrefix:  "",
	})
	user_repo := mocks.NewFakeUserRepository()

	bus := messagebus.NewMessageBus(db)
	bus.Register(domain.CreateUserCommand{}, handlers.NewCreateUserHandler(user_repo))

	return bus
}
