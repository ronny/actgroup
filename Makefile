VERSION ?= dev-$(shell git rev-parse --short HEAD)
BIN_DIR ?= ./bin/

actgroup:
	go build -ldflags="-X main.Version=${VERSION}" -trimpath -o ${BIN_DIR} ./cmd/actgroup
