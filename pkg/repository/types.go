package repository

import (
	"context"
	"io"
)

// User entity
type User struct {
	ID       int
	FullName string
	Email    string
	Password string
}

type storage interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context) error
}

// Vault persistent database
type Vault interface {
	storage

	CreateUser(context.Context, User) error
	UserStream(context.Context, int) (io.ReadCloser, error)
	UserWithOffset(context.Context, int, int) ([]User, error)
	DeleteUserByID(context.Context, int) error
	DeleteUserByEmail(context.Context, string) error
}

// Token entity
type Token string

// Cache temporary storage
type Cache interface {
	storage

	AddToken(context.Context, Token, int) error
	HasToken(context.Context, Token) bool
	DeleteToken(context.Context, Token) error
	UserWithOffset(context.Context, int, int) ([]User, error)
}
