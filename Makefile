APP_NAME = server
BIN_DIR = bin

.PHONY: run build init swag docker

init:
	@go install github.com/swaggo/swag/cmd/swag@latest

swag:init
	@swag init -g ./cmd/main.go -o ./docs --parseDependency

run:swag
	@go run ./cmd/.

build:
	@go build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/.
	@cp config.yaml $(BIN_DIR)/config.yaml

docker:
	@docker build -t $(APP_NAME) .

