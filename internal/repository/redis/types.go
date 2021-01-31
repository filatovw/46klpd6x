package redis

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config - redis connection parameters
type Config struct {
	Host         string
	Port         int
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

// Addr return correct Address
func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// LoadConfig creates config and fill it from environment
func LoadConfig() (*Config, error) {
	config := &Config{
		Host:         "127.0.0.1",
		Port:         6379,
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	}
	if val, ok := os.LookupEnv("REDIS_HOST"); ok {
		config.Host = val
	}
	if val, ok := os.LookupEnv("PORT_PORT"); ok {
		port, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PORT_PORT: %w", err)
		}
		config.Port = port
	}
	if val, ok := os.LookupEnv("REDIS_PASSWORD"); ok {
		config.Password = val
	}
	if val, ok := os.LookupEnv("REDIS_DB"); ok {
		db, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("failed to parse REDIS_DB: %w", err)
		}
		config.DB = db
	}
	return config, nil
}

// User entity
type User struct {
	ID       int
	FullName string
	Email    string
}

// Token entity
type Token string

// Repository works with postgres database
type Repository interface {
	Connect() error
	Disconnect() error
	Ping() error

	AddToken(Token, int) error
	HasToken(Token) bool
	DeleteToken(Token) error
}
