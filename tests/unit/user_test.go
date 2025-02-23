package unit

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForNewUser(t *testing.T) {
	// Arrange
	Name := "ali"
	Age := 20

	//Act
	user, err := entities.NewUser(Name, Age)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, user.UserName, Name)
	assert.Equal(t, user.Age, Age)
}

func TestUserIsUnder18YearsOld(t *testing.T) {
	// Arrange
	errorExpected := errors.BadRequest("User.AgeInvalid")

	//Act
	_, err := UserCreationMethod("", 17)

	//Assert
	assert.Equal(t, err, errorExpected)
}

func TestUserNameIsInvalid(t *testing.T) {
	errorExpected := errors.BadRequest("User.Invalid")

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
	return entities.NewUser(userName, age)
}
