package main

import (
	"log"

	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/config"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/ramDB"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	db, err := ramDB.New()
	if err != nil {
		log.Fatal("db err", err)
	}

	parser, err := currancy.New(config.Currency)
	if err != nil {
		log.Fatal("parser init failed:", err)
	}
	valuteChannel, err := parser.ParseCurrencies()
	if err != nil {
		log.Fatal("valute channel return error:", err)
	}

	tgServer, err := server.New(db, config.TgServer)
	if err != nil {
		log.Fatal("tg server init failed:", err)
	}
	tgServer.InitCurrancies(valuteChannel)

	tgClient, err := client.New(config.TgClient, parser)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient, tgServer)

	tgClient.ListenUpdates(msgModel)
}
