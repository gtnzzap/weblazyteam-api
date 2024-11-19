package main

import (
	"log"
	"net/http"
	"weblazyteam-api/database"
	"weblazyteam-api/routers"
)

func main() {
	// Инициализация базы данных
	database.Connect()

	// Настройка маршрутов
	router := routers.SetupRouter()

	// Запуск сервера
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
