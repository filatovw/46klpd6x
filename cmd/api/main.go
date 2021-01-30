package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/filatovw/46klpd6x/app/api"
	"github.com/filatovw/46klpd6x/app/api/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("can't init logger: %s", err)
		os.Exit(1)
	}
	defer logger.Sync()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	sugar := logger.Sugar()

	if err := godotenv.Load(); err != nil {
		sugar.Infof(".env file was not found: %s", err)
	}

	ctx := context.Background()
	ctx, cancelFn := context.WithCancel(ctx)

	config, err := config.Load()
	if err != nil {
		sugar.Fatalf("failed to get config for application")
	}
	api := api.New(ctx, sugar, config)
	go func() {
		if err := api.Serve(); err != nil {
			os.Exit(1)
		}
	}()

	go func() {
		<-sigC
		sugar.Infow("shutdown the App", "stop this shit", time.Now().Format(time.RFC822Z))
		cancelFn()
	}()
	api.Shutdown()
}
