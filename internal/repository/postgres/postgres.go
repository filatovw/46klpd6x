package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres structure
type Postgres struct {
	context context.Context
	logger  *zap.SugaredLogger
	config  *Config
	conn    *sql.DB
	db      *gorm.DB
}

// Connect to the Postgres database
func (p *Postgres) Connect() error {
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
func (p *Postgres) Disconnect() error {
	return p.conn.Close()
}

// Ping database
func (p *Postgres) Ping() error {
	return p.conn.Ping()
}

// New Postgres instance
func New(ctx context.Context, logger *zap.SugaredLogger, config *Config) Postgres {
	return Postgres{
		context: ctx,
		logger:  logger,
		config:  config,
	}
}
