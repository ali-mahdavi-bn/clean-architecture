package integration

import (
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	queries "clean-hex/internal/user_management/service_layer/queries/user"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewGetUser(t *testing.T) {
	command, ctx := CreateUserCommandCreationMethod("NewAli", 0)

	result, err := Bus.Handle(ctx, command)
	newUser, ok := result.(*entities.User)
	user, err := queries.ViewGetUser(Bus.DB, newUser.ID)
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
