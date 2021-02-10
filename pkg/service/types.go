package service

type User struct {
	Email string
}

type UserManager interface {
	CreateUser(User) error
	DeleteUser(User) error
	Users(int, int) ([]User, error)
}

type AuthManager interface {
	SignIn(User) (string, error)
	SignOut(string) error
	FindByToken(string) (*User, error)
	IsAdminUser(User) bool
}
