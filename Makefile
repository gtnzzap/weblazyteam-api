# Переменные
APP_NAME=weblazyteam-api
SRC_DIR=.
TEST_REPORT=test-report.out

# Компиляция
build:
	@echo "Сборка приложения..."
	@go build -o $(APP_NAME)
	@echo "Сборка завершена."

# Запуск
run: build
	@echo "Запуск сервера..."
	@./$(APP_NAME)

# Запуск тестов
test: test-unit test-integration

test-unit:
	@echo "Запуск юнит-тестов..."
	@go test $(SRC_DIR)/... -run Unit -v

test-integration:
	@echo "Запуск интеграционных тестов..."
	@go test $(SRC_DIR)/... -run Integration -v

# Тесты с генерацией отчета
test-report:
	@echo "Запуск всех тестов и генерация отчета..."
	@go test $(SRC_DIR)/... -cover -coverprofile=$(TEST_REPORT)
	@go tool cover -html=$(TEST_REPORT) -o coverage.html
	@echo "Отчет о покрытии тестами сгенерирован: coverage.html"

# Линтеры
lint:
	@echo "Запуск линтеров..."
	@golangci-lint run

# Очистка
clean:
	@echo "Очистка старых сборок..."
	@rm -f $(APP_NAME)
	@rm -f $(TEST_REPORT)
	@rm -f coverage.html
	@echo "Очистка завершена."

.PHONY: build run test test-unit test-integration test-report lint clean
