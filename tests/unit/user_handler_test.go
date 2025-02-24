package unit

import (
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/errors"
	"clean-hex/tests/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bus = mocks.MockUserManagementBootstrapTestApp()

func TestAddUser(t *testing.T) {
	command, ctx := CreateUserCommandCreationMethod("NewAli", 0)

	result, err := bus.Handle(ctx, command)
	user, ok := result.(*entities.User)

	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, user.UserName, command.UserName)
	assert.Equal(t, user.Age, command.Age)
}

func TestForUserExisting(t *testing.T) {
	command, ctx := CreateUserCommandCreationMethod("", 0)

	result, err := bus.Handle(ctx, command)

	assert.Equal(t, err, errors.BadRequest("User.AlreadyExists"))
	assert.Nil(t, result)
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
