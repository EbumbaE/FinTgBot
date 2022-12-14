package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/EbumbaE/FinTgBot/bot/internal/clients/producer"
	client "github.com/EbumbaE/FinTgBot/bot/internal/clients/tg"
	"github.com/EbumbaE/FinTgBot/bot/internal/config"
	"github.com/EbumbaE/FinTgBot/bot/internal/currency"
	"github.com/EbumbaE/FinTgBot/bot/internal/model/messages"
	"github.com/EbumbaE/FinTgBot/bot/internal/servers/sender"
	server "github.com/EbumbaE/FinTgBot/bot/internal/servers/tg"
	"github.com/EbumbaE/FinTgBot/bot/internal/storage/psql"
	"github.com/EbumbaE/FinTgBot/bot/pkg/logger"
	"github.com/EbumbaE/FinTgBot/bot/pkg/metrics"
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
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("close database", zap.Error(err))
		}
	}()

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
		logger.Error("init producer", zap.Error(err))
	}
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	producer.StartConsumeError(ctx)
	defer func() {
		if err := producer.Close(); err != nil {
			logger.Error("close producer", zap.Error(err))
		}
	}()

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

	senderServer := sender.New(config.SenderServer, tgClient)
	ctx.Value("allDoneWG").(*sync.WaitGroup).Add(1)
	if err := senderServer.StartServe(ctx); err != nil {
		logger.Error("Start Serve", zap.Error(err))
	}

	ctx.Value("allDoneWG").(*sync.WaitGroup).Wait()
	logger.Info("All is shutdown")
}
