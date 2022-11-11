package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/producer"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
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

	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	metrics.InitServer(ctx)

	config, err := config.New()
	if err != nil {
		logger.Fatal("config init failed: ", zap.Error(err))
	}

	db, err := psql.New(config.PsqlDatabase)
	if err != nil {
		logger.Fatal("db init: ", zap.Error(err))
	}
	if err := db.CheckHealth(); err != nil {
		logger.Fatal("db check health: ", zap.Error(err))
	}

	parser, err := currency.New(config.Currency)
	if err != nil {
		logger.Fatal("parser init failed:", zap.Error(err))
	}
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	err = parser.ParseCurrencies(ctx, db)
	if err != nil {
		logger.Fatal("valute channel return error:", zap.Error(err))
	}

	producer := producer.New(config.Producer)
	if err := producer.InitProducer(ctx); err != nil {
		logger.Fatal("init producer", zap.Error(err))
	}
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	producer.StartConsumeError()

	tgServer, err := server.New(db, producer, config.TgServer)
	if err != nil {
		logger.Fatal("tg server init failed:", zap.Error(err))
	}
	tgServer.InitMiddleware()

	tgClient, err := client.New(config.TgClient, parser)
	if err != nil {
		logger.Fatal("tg client init failed:", zap.Error(err))
	}
	tgClient.InitMiddleware()

	msgModel := messages.New(tgClient, tgServer)
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	tgClient.ListenUpdates(ctx, msgModel)

	ctx.Value("allDoneWG").(*sync.WaitGroup).Wait()
	if err := db.Close(); err != nil {
		logger.Error("Close database", zap.Error(err))
	}
	logger.Info("All is shutdown")
}
