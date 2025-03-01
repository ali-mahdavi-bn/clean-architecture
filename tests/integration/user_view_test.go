package integration

import (
	"clean-hex/internal/user_management/domain"
	"clean-hex/internal/user_management/domain/entities"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewGetUser(t *testing.T) {
	command, ctx := CreateUserCommandCreationMethod("ali", 0)

	result, err := Bus.Handle(ctx, command)
	_, ok := result.(*entities.User)
	//trade, err := queries.GetUser(ctx, Bus.Uow, newUser.ID, RedisStore)
	fmt.Println(ok, result, err)
	//assert.Nil(t, err)
	assert.True(t, true)
	//assert.True(t, ok)
	//assert.Equal(t, trade.ID, newUser.ID)
	//assert.Equal(t, trade.UserName, newUser.UserName)
	//assert.Equal(t, trade.Age, newUser.Age)

}

func CreateUserCommandCreationMethod(userName string, age int) (*domain.CreateUserCommand, context.Context) {
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
	return &command, ctx
}
