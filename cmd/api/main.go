package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/filatovw/46klpd6x/app/api"
	"github.com/filatovw/46klpd6x/app/api/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// setup logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("can't init logger: %s", err)
		os.Exit(1)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// handle Ctrl-C
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	// load .env vars
	if err := godotenv.Load(); err != nil {
		sugar.Infof(".env file was not found: %s", err)
	}

	ctx := context.Background()
	ctx, cancelFn := context.WithCancel(ctx)

	// load api config
	config, err := config.Load()
	if err != nil {
		sugar.Fatalf("failed to get config for application")
	}

	configPG, _ := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	conn := stdlib.OpenDB(*configPG)
	defer conn.Close()
	gormConn, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	res := ""
	gormConn.Raw("select 1").Find(&res)
	sugar.Infof("HELLO %+v\n", res)

	// create api server
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
