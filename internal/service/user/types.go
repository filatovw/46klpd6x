package user

import (
	"github.com/filatovw/46klpd6x/internal/repository/postgres"
	"github.com/filatovw/46klpd6x/internal/repository/redis"
	"go.uber.org/zap"
)

type User struct {
}

type Service interface {
	CreateUser(User) error
	DeleteUser(User) error
	Users(int, int) ([]User, error)
}

type UserService struct {
	logger *zap.SugaredLogger
	db     postgres.Repository
	cache  redis.Repository
}

func New(logger *zap.SugaredLogger, db postgres.Repository, cache redis.Repository) UserService {
	return UserService{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (s *UserService) CreateUser(user User) error                  { return nil }
func (s *UserService) DeleteUser(user User) error                  { return nil }
func (s *UserService) Users(limit int, offset int) ([]User, error) { return []User{}, nil }
