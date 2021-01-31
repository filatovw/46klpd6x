package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/filatovw/46klpd6x/app/api"
	"github.com/filatovw/46klpd6x/internal/repository/postgres"
	"github.com/filatovw/46klpd6x/internal/repository/redis"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// load .env vars
	godotenv.Load()
	// load api config
	apiConfig, err := api.LoadConfig()
	if err != nil {
		fmt.Printf("failed to get config for application: %s", err)
		os.Exit(1)
	}

	// setup logger
	var logger *zap.Logger
	if apiConfig.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		fmt.Printf("can't init logger: %s", err)
		os.Exit(1)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// handle Ctrl-C
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	ctx := context.Background()
	ctx, cancelFn := context.WithCancel(ctx)

	// connect to Postgres
	pgConfig, err := postgres.LoadConfig()
	if err != nil {
		sugar.Fatalf("failed to get config for Postgres")
	}
	pg := postgres.New(sugar, pgConfig)
	if err := pg.Connect(); err != nil {
		sugar.Fatalf("failed to connect to the Postgres instance: %s", err)
	}
	defer pg.Disconnect()
	if err := pg.Ping(); err != nil {
		sugar.Fatalf("Unable to Ping DB: %s", err)
	}

	// connect to Redis
	redisConfig, err := redis.LoadConfig()
	if err != nil {
		sugar.Fatalf("failed to get config for Redis")
	}
	redis := redis.New(sugar, *redisConfig)
	redis.Connect()
	if err := redis.Ping(ctx); err != nil {
		sugar.Fatalf("Unable to Ping DB: %s", err)
	}
	defer redis.Disconnect()

	// create api server
	api := api.New(ctx, sugar, apiConfig)
	go func() {
		if err := api.Serve(); err != nil {
			sugar.Warnf("api service stopped: %s", err)
		}
	}()

	go func() {
		<-sigC
		sugar.Infow("shutdown the App", "stopped at", time.Now().Format(time.RFC822Z))
		cancelFn()
	}()
	api.Shutdown()
}
