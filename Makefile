# Укажите имя вашего основного файла
APP_NAME=weblazyteam-api
MAIN_FILE=main.go

# Команды для работы с сервером
.PHONY: run
run: ## Запуск сервера
	@echo "Запуск сервера..."
	go run $(MAIN_FILE)

.PHONY: build
build: ## Сборка исполняемого файла
	@echo "Сборка приложения..."
	go build -o $(APP_NAME)

.PHONY: start
start: build ## Сборка и запуск исполняемого файла
	@echo "Запуск собранного приложения..."
	./$(APP_NAME)

.PHONY: clean
clean: ## Очистка собранного файла
	@echo "Очистка проекта..."
	rm -f $(APP_NAME)

# Команды для проверки и тестирования
.PHONY: test
test: ## Запуск тестов
	@echo "Запуск тестов..."
	go test ./... -v

.PHONY: lint
lint: ## Запуск линтера
	@echo "Проверка кода линтером..."
	golangci-lint run

.PHONY: fmt
fmt: ## Форматирование кода
	@echo "Форматирование Go-кода..."
	go fmt ./...

.PHONY: vet
vet: ## Проверка кода на ошибки
	@echo "Анализ Go-кода..."
	go vet ./...

.PHONY: check
check: fmt vet lint ## Форматирование, линтинг и анализ кода
	@echo "Проверка кода завершена."

.PHONY: help
help: ## Вывод доступных команд
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
