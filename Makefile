APP_NAME = server
BIN_DIR = bin

init:
	@go install github.com/swaggo/swag/cmd/swag@latest

swag:init
	@swag init -g ./cmd/main.go -o ./docs --parseDependency

run:
	@go run ./cmd/.

build:
	@go build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/.

.PHONY: run build init swag
