package integration

import (
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	queries "clean-hex/internal/user_management/service_layer/queries/user"
	"clean-hex/tests"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bus = tests.MockUserManagementBootstrapTestApp()

func TestViewGetUser(t *testing.T) {
	command, ctx := CreateUserCommandCreationMethod("", 0)

	result, err := bus.Handle(ctx, command)
	newUser, ok := result.(*entities.User)
	user, err := queries.ViewGetUser(bus.DB, newUser.ID)
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, user.ID, newUser.ID)
	assert.Equal(t, user.UserName, newUser.UserName)
	assert.Equal(t, user.Age, newUser.Age)

}

func CreateUserCommandCreationMethod(userName string, age int) (domain.CreateUserCommand, context.Context) {
	if userName == "" {
		userName = "ali"
	}
	if age == 0 {
		age = 20
	}
	ctx := context.Background()
	command := domain.CreateUserCommand{
		UserName: userName,
		Age:      age,
	}
	return command, ctx
}
