package entities

import (
	"clean-hex/pkg/framwork/adapter"
	"clean-hex/pkg/framwork/errors"
)

type User struct {
	adapter.BaseEntity
	Age      int
	UserName string
	Amount   int
	Trades   []Trade
}

func NewUser(userName string, age int, amount int) (*User, error) {
	if userName == "admin" {
		return nil, errors.BadRequest("Transaction.Invalid")
	}
	if age < 18 {
		return nil, errors.BadRequest("Transaction.AgeInvalid")
	}
	user := &User{}
	user.UserName = userName
	user.Age = age
	user.Amount = amount
	return user, nil
}

func (u *User) Update(userName string, age int, amount int) error {
	if userName == "admin" {
		return errors.BadRequest("Transaction.Invalid")
	}
	if age < 18 {
		return errors.BadRequest("Transaction.AgeInvalid")
	}

	u.UserName = userName
	u.Age = age
	u.Amount = amount
	return nil
}
