package main

import (
	"context"
	"fmt"
	"log"
	"os"

	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/ramDB"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	file, err := os.OpenFile("l.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	db, err := ramDB.New()
	if err != nil {
		log.Fatal("db err", err)
	}

	parser, err := currency.New(config.Currency)
	if err != nil {
		log.Fatal("parser init failed:", err)
	}
	rateCurrencyChan, err := parser.ParseCurrencies(ctx)
	if err != nil {
		log.Fatal("valute channel return error:", err)
	}

	tgServer, err := server.New(db, config.TgServer)
	if err != nil {
		log.Fatal("tg server init failed:", err)
	}
	tgServer.InitCurrancies(ctx, rateCurrencyChan)

	tgClient, err := client.New(config.TgClient, parser)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient, tgServer)
	tgClient.ListenUpdates(ctx, msgModel)

	var command string
	for {
		fmt.Scanln(&command)
		if command == "shutdown" {
			cancel()
			break
		}
	}
}
