package unit

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/framwork/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForNewUser(t *testing.T) {
	// Arrange
	name := "ali"
	age := 20
	amount := 20

	//Act
	user, err := entities.NewUser(name, age, amount)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, user.UserName, name)
	assert.Equal(t, user.Age, age)
	assert.Equal(t, user.Amount, amount)
}

func TestUserIsUnder18YearsOld(t *testing.T) {
	// Arrange
	errorExpected := errors.BadRequest("Transaction.AgeInvalid")

	//Act
	_, err := UserCreationMethod("", 17)

	//Assert
	assert.Equal(t, err, errorExpected)
}

func TestUserNameIsInvalid(t *testing.T) {
	errorExpected := errors.BadRequest("Transaction.Invalid")

	_, err := UserCreationMethod("admin", 0)

	assert.Equal(t, err, errorExpected)
}

func UserCreationMethod(userName string, age int) (*entities.User, error) {
	if userName == "" {
		userName = "ali"
	}
	if age == 0 {
		age = 20
	}
	amount := 20
	return entities.NewUser(userName, age, amount)
}
