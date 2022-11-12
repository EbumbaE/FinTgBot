protoc -I. \
  -I"${GOPATH}"/src \
  -I"${GOPATH}"/src/gitlab.ozon.ru/agarkov/route256/lectures/4-class/gRPC/metadata/third_party \
  --go_out=. --go_opt=paths=source_relative echo.proto
