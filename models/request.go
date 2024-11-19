package models

// Структура заявки
type Request struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Email       string `json:"email,omitempty"` // Email может быть пустым
	RequestText string `json:"request_text"`
}
