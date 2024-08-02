# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifeq ($(OS),Windows_NT)
	BINARY = bin/$(APP_NAME).exe
else
	BINARY = bin/$(APP_NAME)
endif

MAIN=cmd/main.go

MIGRATION_DIR="migrations"
MIGRATION_CMD=@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_DSN) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR)

dev:
	@go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build --tags dev -o ${BINARY} ${MAIN}" --build.bin "${BINARY}" --build.delay "100" \
	--build.exclude_dir "" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \
	--screen.clear_on_rebuild true \
	--log.main_only true

swag:
	@go run github.com/swaggo/swag/cmd/swag@latest init --parseDependency

db-up:
	$(MIGRATION_CMD) up

db-down:
	$(MIGRATION_CMD) down

db-status:
	$(MIGRATION_CMD) status

db-reset:
	$(MIGRATION_CMD) reset

db-new:
	$(MIGRATION_CMD) create $(filter-out $@,$(MAKECMDGOALS)) sql