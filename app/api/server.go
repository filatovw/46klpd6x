package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/filatovw/46klpd6x/app/api/config"
	"go.uber.org/zap"
)

// API HTTP server
type API struct {
	ctx    context.Context
	server *http.Server
	logger *zap.SugaredLogger
	config *config.Config
}

// Serve start server
func (a *API) Serve() error {
	a.logger.Infow("starting server", "connection_string", a.config.ConnectionString())
	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// Shutdown server
func (a *API) Shutdown() error {
	<-a.ctx.Done()
	return a.server.Shutdown(a.ctx)
}

// New - instantiate API HTTP server
func New(ctx context.Context, logger *zap.SugaredLogger, config *config.Config) API {
	server := &http.Server{
		Addr:           config.ConnectionString(),
		Handler:        routes(),
		ReadTimeout:    time.Duration(config.ReadTimeout),
		WriteTimeout:   time.Duration(config.WriteTimeout),
		MaxHeaderBytes: config.MaxHeaderBytes,
	}
	return API{
		ctx:    ctx,
		server: server,
		logger: logger,
		config: config,
	}
}
