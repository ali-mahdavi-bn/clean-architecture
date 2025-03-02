package entities

import "clean-hex/pkg/framwork/adapter"

type Trade struct {
	adapter.BaseEntity
	UserID uint   `gorm:"index"`
	Stock  string `gorm:"not null"`
	Price  int    `gorm:"not null"`
	Amount int    `gorm:"not null"`
}

func NewTrade(userID uint, stock string, price int, amount int) (*Trade, error) {
	trade := &Trade{}
	trade.UserID = userID
	trade.Stock = stock
	trade.Price = price
	trade.Amount = amount
	return trade, nil
}

func (u *Trade) Update(userID uint, stock string, price int, amount int) error {
	u.UserID = userID
	u.Stock = stock
	u.Price = price
	u.Amount = amount
	return nil
}
