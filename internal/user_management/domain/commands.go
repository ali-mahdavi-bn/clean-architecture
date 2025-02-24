package domain

type CreateUserCommand struct {
	UserName string
	Age      int
}

type UpdateUserCommand struct {
	UserName string
	Age      int
	UserId   uint
}

type DeleteUserCommand struct {
	UserId uint
}
