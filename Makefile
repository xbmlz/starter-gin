APP_NAME = server
BIN_DIR = bin

run:
	@go run ./cmd/.

build:
	@go build -ldflags="-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/.

.PHONY: run build
