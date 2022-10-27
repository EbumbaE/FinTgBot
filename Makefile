PKG := "gitlab.ozon.dev/ivan.hom.200/telegram-bot/cmd/bot"
BIN := "bin/bot"

all: build

build: 
	go build -o ${BIN}/bot.exe ${PKG}

test:
	go test ./...

dev:
	go run workshop5/cmd/fibonacci -devel

prod:
	go run ${PKG} 2>&1 | tee logger/log.txt

logs:
	cd logger && docker compose up

tracing:
	cd tracing && docker compose up

metrics:
	cd metrics && docker compose up

pull:
	docker pull prom/prometheus
	docker pull grafana/grafana-oss
	docker pull ozonru/file.d:latest-linux-amd64
	docker pull elasticsearch:7.17.6
	docker pull graylog/graylog:4.3
	docker pull jaegertracing/all-in-one:1.18
