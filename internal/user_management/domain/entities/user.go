package entities

import (
	"clean-hex/pkg/errors"
	"clean-hex/pkg/framwork/adapter"
)

type User struct {
	adapter.BaseEntity
	Age      int
	UserName string
}

func NewUser(UserName string, Age int) (*User, error) {
	if UserName == "admin" {
		return nil, errors.BadRequest("User.Invalid")
	}
	if Age < 18 {
		return nil, errors.BadRequest("User.AgeInvalid")
	}
	user := &User{}
	user.UserName = UserName
	user.Age = Age
	return user, nil
}

func (u *User) Update(UserName string, Age int) error {
	if UserName == "admin" {
		return errors.BadRequest("User.Invalid")
	}
	if Age < 18 {
		return errors.BadRequest("User.AgeInvalid")
	}

	u.UserName = UserName
	u.Age = Age
	return nil
}
