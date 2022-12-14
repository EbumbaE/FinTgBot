PROJECT := github.com/EbumbaE/FinTgBot
PKG_BOT := ${PROJECT}/bot/cmd/bot
PKG_REPORT := ${PROJECT}/report/cmd/report

all: build

build: 
	cd bot && go build -o bin/bot/bot.exe ${PKG_BOT}	
	cd report && go build -o bin/report/report.exe ${PKG_REPORT}

test:
	cd bot && go test ./...
	cd report && go test ./...

logs:
	cd pkg/logs && docker compose up

tracing:
	cd pkg/tracing && docker compose up

metrics:
	cd pkg/metrics && docker compose up

cache:
	cd pkg/cache && docker compose up

kafka: 
	cd pkg/kafka && docker compose up

database:
	cd pkg/database && docker compose up
	goose postgres "user=postgres password=1234 dbname=postgres sslmode=disable" up

pull:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	docker pull prom/prometheus
	docker pull grafana/grafana-oss
	docker pull ozonru/file.d:latest-linux-amd64
	docker pull elasticsearch:7.17.6
	docker pull graylog/graylog:4.3
	docker pull jaegertracing/all-in-one:1.18
	docker pull memcached
	docker pull wurstmeister/kafka
	docker pull postgres