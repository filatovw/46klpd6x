package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"io"

	"github.com/filatovw/46klpd6x/pkg/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres structure
type Postgres struct {
	logger *zap.SugaredLogger
	config *Config
	conn   *sql.DB
	db     *gorm.DB
}

// Connect to the Postgres database
func (p *Postgres) Connect(_ context.Context) error {
	pgxConfig, err := pgx.ParseConfig(p.config.ConnectionString)
	if err != nil {
		return fmt.Errorf("unable to parse postgres connection string: %s", err)
	}
	conn := stdlib.OpenDB(*pgxConfig)
	gormConn, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open Postgres connection: %w", err)
	}
	p.db = gormConn
	p.conn = conn
	return nil
}

// Disconnect database
func (p *Postgres) Disconnect(_ context.Context) error {
	return p.conn.Close()
}

// Ping database
func (p *Postgres) Ping(_ context.Context) error {
	return p.conn.Ping()
}
func (p *Postgres) CreateUser(ctx context.Context, user repository.User) error { return nil }
func (p *Postgres) UserStream(ctx context.Context, offset int) (io.ReadCloser, error) {
	return nil, nil
}
func (p *Postgres) UserWithOffset(ctx context.Context, offset int, limit int) ([]repository.User, error) {
	return []repository.User{}, nil
}
func (p *Postgres) DeleteUserByID(ctx context.Context, id int) error          { return nil }
func (p *Postgres) DeleteUserByEmail(ctx context.Context, email string) error { return nil }

// New Postgres instance
func New(logger *zap.SugaredLogger, config *Config) Postgres {
	return Postgres{
		logger: logger,
		config: config,
	}
}
