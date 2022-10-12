package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/psql"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	file, err := os.OpenFile("l.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}

	db, err := psql.New(config.PsqlDatabase)
	if err != nil {
		log.Fatal("db init: ", err)
	}
	if err := db.CheckHealth(); err != nil {
		log.Fatal("db check health: ", err)
	}

	parser, err := currency.New(config.Currency)
	if err != nil {
		log.Fatal("parser init failed:", err)
	}
	err = parser.ParseCurrencies(ctx, db)
	if err != nil {
		log.Fatal("valute channel return error:", err)
	}

	tgServer, err := server.New(db, config.TgServer)
	if err != nil {
		log.Fatal("tg server init failed:", err)
	}

	tgClient, err := client.New(config.TgClient, parser)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient, tgServer)
	tgClient.ListenUpdates(ctx, msgModel)
}
