package auth

import (
	"strings"

	"github.com/filatovw/46klpd6x/internal/repository/postgres"
	"github.com/filatovw/46klpd6x/internal/repository/redis"
	"go.uber.org/zap"
)

// User model in Auth service
type User struct {
	Email string
}

// Service auth user, check permission
type Service interface {
	SignIn(User) (string, error)
	SignOut(string) error
	FindByToken(string) (*User, error)
	IsAdminUser(User) bool
}

// AuthService implementation of auth service
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

func (s *AuthService) SignIn(user User) (string, error)        { return "", nil }
func (s *AuthService) SignOut(token string) error              { return nil }
func (s *AuthService) FindByToken(token string) (*User, error) { return nil, nil }
func (s *AuthService) IsAdminUser(user User) bool {
	return strings.HasSuffix(user.Email, "@test.com")
}
