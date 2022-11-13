package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/sender"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/consumer"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/cache"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/psql"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/logger"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/pkg/metrics"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "allDoneWG", &sync.WaitGroup{})
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	config, err := config.New()
	if err != nil {
		logger.Fatal("config init failed", zap.Error(err))
	}

	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	metrics.InitServer(ctx)

	db, err := psql.New(config.PsqlDatabase)
	if err != nil {
		logger.Fatal("db init: ", zap.Error(err))
	}
	if err := db.CheckHealth(); err != nil {
		logger.Fatal("db check health: ", zap.Error(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("close database", zap.Error(err))
		}
	}()

	cache := cache.New("127.0.0.1:11211")
	if err := cache.Ping(); err != nil {
		logger.Error("cache ping: ", zap.Error(err))
	}

	senderClient := sender.New(config.SenderClient)
	if err := senderClient.StartConnection(); err != nil {
		logger.Error("start connection in sender client", zap.Error(err))
	}
	defer func() {
		if err := senderClient.Close(); err != nil {
			logger.Error("close connection in sender client", zap.Error(err))
		}
	}()

	reportConsumer := consumer.New(config.Consumer, db, cache, senderClient)
	if err := reportConsumer.InitConsumerGroup(ctx); err != nil {
		logger.Fatal("Init Consume group", zap.Error(err))
	}
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	reportConsumer.StartConsumerGroup(ctx)

	ctx.Value("allDoneWG").(*sync.WaitGroup).Wait()
	logger.Info("All is shutdown")
}
