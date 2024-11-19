package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"weblazyteam-api/database"
	"weblazyteam-api/models"
	"weblazyteam-api/utils"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Регистрация пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Декодирование JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Проверка на валидность входных данных
	if err := validate.Struct(user); err != nil {
		http.Error(w, "Некорректные данные: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Хэширование пароля
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Вставка нового пользователя в базу данных
	query := "INSERT INTO users (login, password, name, role) VALUES (@p1, @p2, @p3, @p4)"
	_, err = database.DB.Exec(query, user.Login, hashedPassword, user.Name, "user") // Новые пользователи получают роль "user" по умолчанию
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// Авторизация пользователя
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds models.User

	// Декодирование JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Получение данных пользователя из БД
	var storedPassword, role string
	query := "SELECT password, role FROM users WHERE login = @p1"
	err := database.DB.QueryRow(query, creds.Login).Scan(&storedPassword, &role)
	if err != nil {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	// Проверка пароля
	if !utils.CheckPasswordHash(creds.Password, storedPassword) {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		return
	}

	// Генерация JWT
	token, err := utils.GenerateJWT(creds.Login, role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Установка токена в cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}

// Получение всех пользователей (для разработчиков и администраторов)
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, login, name, role FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Login, &user.Name, &user.Role)
		if err != nil {
			http.Error(w, "Error reading user data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Изменение роли пользователя (только для администраторов)
func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	type RoleUpdate struct {
		UserID int    `json:"user_id"`
		Role   string `json:"role"`
	}

	var update RoleUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET role = @p1 WHERE id = @p2"
	_, err := database.DB.Exec(query, update.Role, update.UserID)
	if err != nil {
		http.Error(w, "Failed to update user role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role updated successfully"))
}
