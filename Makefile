PKG := "gitlab.ozon.dev/ivan.hom.200/telegram-bot/cmd/bot"
BIN := "bin/bot"

all: build

build: 
	go build -o ${BIN}/bot.exe ${PKG}

test:
	go test ./...