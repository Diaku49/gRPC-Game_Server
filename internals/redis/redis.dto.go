package redis

type CreateUserDTO struct {
	Name  string
	Email string
}

type GetUser struct {
	Name  string
	Email string
}
