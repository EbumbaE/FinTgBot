PROJECT := gitlab.ozon.dev/ivan.hom.200/telegram-bot
PKG_BOT := ${PROJECT}/cmd/bot
PKG_REPORT := ${PROJECT}/cmd/report

all: build

build: 
	cd bot && go build -o bin/bot/bot.exe ${PKG_BOT}	
	cd report && go build -o bin/report/report.exe ${PKG_REPORT}

test:
	cd bot && go test ./...
	cd report && go test ./...

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