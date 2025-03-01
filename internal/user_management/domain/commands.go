package domain

// trade
type CreateUserCommand struct {
	UserName string
	Age      int
	Amount   int
}
type UpdateUserCommand struct {
	UserName string
	Age      int
	Amount   int
	UserId   uint
}
type DeleteUserCommand struct {
	UserId uint
}

// user
type CreateTradeCommand struct {
	UserId uint
	Stock  string
	Price  int
	Amount int
}
type UpdateTradeCommand struct {
	UserId uint
	Stock  string
	Price  int
	Amount int
}
type DeleteTradeCommand struct {
	TradeId uint
}
