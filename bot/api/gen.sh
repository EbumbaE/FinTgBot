protoc -I. \
  -I"${GOPATH}"/src \
  -I"${GOPATH}"/src/gitlab.ozon.dev/ivan.hom.200/telegram-bot/bot/api \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  --go_out=. --go_opt=paths=source_relative \
  api.proto 
