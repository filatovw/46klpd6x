package auth

import (
	"strings"

	"github.com/filatovw/46klpd6x/pkg/repository"
	"github.com/filatovw/46klpd6x/pkg/service"
	"go.uber.org/zap"
)

// Service implementation of auth service
type Service struct {
	logger *zap.SugaredLogger
	db     repository.Vault
	cache  repository.Cache
}

// New auth service
func New(logger *zap.SugaredLogger, db repository.Vault, cache repository.Cache) Service {
	return Service{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (s *Service) SignIn(user service.User) (string, error)        { return "", nil }
func (s *Service) SignOut(token string) error                      { return nil }
func (s *Service) FindByToken(token string) (*service.User, error) { return nil, nil }
func (s *Service) IsAdminUser(user service.User) bool {
	return strings.HasSuffix(user.Email, "@test.com")
}
