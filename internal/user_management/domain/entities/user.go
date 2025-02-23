package entities

import (
	"clean-hex/pkg/errors"
)

type User struct {
	UserName string
	Age      int
	ID       uint `gorm:"primaryKey"`
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

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) Update(Name string, Age int) (*User, error) {
	u.UserName = Name
	if Age < 18 {
		return nil, errors.BadRequest("User.AgeInvalid")
	}
	u.Age = Age
	return u, nil
}
