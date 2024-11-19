package models

// Структура пользователя
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password,omitempty"` // Пароль не возвращается в ответе
	Name     string `json:"name"`
	Role     string `json:"role"`
}
