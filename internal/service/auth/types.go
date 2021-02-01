package auth

import (
	"github.com/filatovw/46klpd6x/internal/repository/postgres"
	"github.com/filatovw/46klpd6x/internal/repository/redis"
	"go.uber.org/zap"
)

type User struct{}

type Service interface {
	SignIn(User) error
	SignOut(string) error
}

type AuthService struct {
	logger *zap.SugaredLogger
	db     postgres.Repository
	cache  redis.Repository
}

func New(logger *zap.SugaredLogger, db postgres.Repository, cache redis.Repository) AuthService {
	return AuthService{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (s *AuthService) SignIn(user User) error     { return nil }
func (s *AuthService) SignOut(token string) error { return nil }
