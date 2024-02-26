APP_NAME = starter-gin
BIN_DIR = bin

run:
	@go run ./cmd/main.go

build:
	@go build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/main.go

.PHONY: run build