PROTO_DIRS = proto
PROTO_FILES := $(shell find $(PROTO_DIRS) -type f -name '*.proto')

.PHONY: build test lint run migrup migrdown migrcreate swago
build: ## Собираем бинарник
	CGO_ENABLED=0 go build -o bin/app -mod=readonly ./cmd/app
test: ## Запускаем тесты
	go test -v ./... -race -cover

lint: ## Запускаем линтер (конфиг .golangci.yaml)
	golangci-lint run

run: ## Запускаем сервис (без билда)
	go run ./cmd/app/main.go

swago: ## Запускаем swag (генерим сваггер)
	# https://github.com/swaggo/swag
	# go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency --parseInternal --propertyStrategy pascalcase --parseDepth 3 -g cmd/app/app.go

protogen: ## Генерим proto файлы проекта
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_FILES)

generate: ## Запускаем весь автоген проекта
	@make swago
	@make protogen

migrup: ## Накатываем миграции
#	@make gooseinstall
	go run ./cmd/migrator/migrator.go up

migrdown: ## Откатываем миграции
	@make gooseinstall
	go run ./cmd/migrator/migrator.go down

gooseinstall:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrcreate: ## Создаёт новую миграцию (параметр ARG1 обязателен)
	@if [ -z "$(ARG1)" ]; then \
		echo "Ошибка: Не задан параметр ARG1! Пример: make migrcreate ARG1=my_migration"; \
	else \
		go run ./cmd/migrator/migrator.go create $(ARG1) sql; \
	fi
help:  ## Показываем список команд
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
