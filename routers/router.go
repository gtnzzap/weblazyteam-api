package routers

import (
	"weblazyteam-api/handlers"
	"weblazyteam-api/middleware"

	"github.com/gorilla/mux"
)

// Настройка маршрутов
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Маршруты для пользователей
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Защищённые маршруты
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuth)

	// Роуты для заявок
	api.HandleFunc("/requests", handlers.CreateRequest).Methods("POST")                                // Создание заявки
	api.HandleFunc("/requests", handlers.GetAllRequests).Methods("GET").Queries("role", "developer")   // Получение всех заявок
	api.HandleFunc("/requests", handlers.DeleteRequest).Methods("DELETE").Queries("role", "developer") // Удаление заявки

	// Роуты для пользователей (админ/разработчик)
	api.HandleFunc("/users", handlers.GetAllUsers).Methods("GET").Queries("role", "developer")       // Получение всех пользователей
	api.HandleFunc("/users/role", handlers.UpdateUserRole).Methods("PATCH").Queries("role", "admin") // Изменение роли

	return router
}
