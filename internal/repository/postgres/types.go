package postgres

import (
	"io"
	"os"
	"time"
)

// Config - postgres connection parameters
type Config struct {
	ConnectionString string
}

// LoadConfig creates config and fill it from environment
func LoadConfig() (*Config, error) {
	config := &Config{}
	if val, ok := os.LookupEnv("DATABASE_URL"); ok {
		config.ConnectionString = val
	}
	return config, nil
}

// User entity
type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Salt      string
	CreatedAt time.Time
	DeletedAt *time.Time
	IsDeleted bool
}

// Repository works with postgres database
type Repository interface {
	Connect() error
	Disconnect() error
	Ping() error
	CreateUser(User) error
	UserStream(int) (io.ReadCloser, error)
	UserWithOffset(int, int) ([]User, error)
	DeleteUserByID(int) error
	DeleteUserByEmail(string) error
}
