package user

import (
	"github.com/filatovw/46klpd6x/pkg/repository"
	"github.com/filatovw/46klpd6x/pkg/service"
	"go.uber.org/zap"
)

// Service that operates with user profiles
type Service struct {
	logger *zap.SugaredLogger
	db     repository.Vault
	cache  repository.Cache
}

// New user service
func New(logger *zap.SugaredLogger, db repository.Vault, cache repository.Cache) Service {
	return Service{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (s *Service) CreateUser(user service.User) error                  { return nil }
func (s *Service) DeleteUser(user service.User) error                  { return nil }
func (s *Service) Users(limit int, offset int) ([]service.User, error) { return []service.User{}, nil }
