PROJECT := gitlab.ozon.dev/ivan.hom.200/telegram-bot
PKG_BOT := ${PROJECT}/bot/cmd/bot
PKG_REPORT := ${PROJECT}/report/cmd/report
MOD_BOT := ${PROJECT}/bot/go.mod
MOD_REPORT := ${PROJECT}/report/go.mod

all: build

build: 
	go build -o bot/bin/bot/bot.exe -modfile ${MOD_BOT} ${PKG_BOT}	
	go build -o report/bin/report/report.exe -modfile ${MOD_REPORT} ${PKG_REPORT}

test:
	go test ./...

logs:
	cd pkg/logger/logs && docker compose up

tracing:
	cd pkg/tracing && docker compose up

metrics:
	cd pkg/metrics && docker compose up

cache:
	cd pkg/cache && docker compose up

kafka: 
	cd pkg/kafka && docker compose up

pull:
	docker pull prom/prometheus
	docker pull grafana/grafana-oss
	docker pull ozonru/file.d:latest-linux-amd64
	docker pull elasticsearch:7.17.6
	docker pull graylog/graylog:4.3
	docker pull jaegertracing/all-in-one:1.18
	docker pull memcached
	docker pull wurstmeister/kafka