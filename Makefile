APP_NAME = starter-gin
BIN_DIR = bin

run:
	@go run ./cmd/.

build:
	@go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/.

.PHONY: run build