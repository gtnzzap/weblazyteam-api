package models

// Структура пользователя
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" validate:"required,min=3,max=50"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=255"` // Пароль не возвращается в ответе
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Role     string `json:"role" validate:"required,oneof=user developer admin"`
}
