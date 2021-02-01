package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Redis repository
type Redis struct {
	logger *zap.SugaredLogger
	client redis.UniversalClient
	config Config
}

// Connect to Redis server
func (r *Redis) Connect() {
	r.client = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{r.config.Addr()},
		Password: r.config.Password,
		DB:       r.config.DB,
	})
}

// Ping Redis server
func (r Redis) Ping(ctx context.Context) error {
	pong, err := r.client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("ping failed: %s", err)
	}
	if pong != "PONG" {
		r.logger.Warnf("unexpected response message: \"%s\"", pong)
	}
	return nil
}

// Disconnect Redis
func (r *Redis) Disconnect() error {
	return r.client.Close()
}

func (r *Redis) AddToken(token Token, ttl int) error { return nil }
func (r Redis) HasToken(token Token) bool            { return false }
func (r *Redis) DeleteToken(token Token) error       { return nil }

// New Redis repository
func New(logger *zap.SugaredLogger, config Config) Redis {
	return Redis{
		logger: logger,
		config: config,
	}
}
