package handlers

import (
	"encoding/json"
	"net/http"
	"weblazyteam-api/database"
	"weblazyteam-api/models"
)

// Создание новой заявки (доступно всем пользователям)
func CreateRequest(w http.ResponseWriter, r *http.Request) {
	var req models.Request

	// Декодирование JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Вставка новой заявки в базу данных
	query := "INSERT INTO requests (phone_number, full_name, email, request_text) VALUES (@p1, @p2, @p3, @p4)"
	_, err := database.DB.Exec(query, req.PhoneNumber, req.FullName, req.Email, req.RequestText)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Request created successfully"))
}

// Получение всех заявок (доступно разработчикам и администраторам)
func GetAllRequests(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, phone_number, full_name, email, request_text FROM requests")
	if err != nil {
		http.Error(w, "Failed to retrieve requests", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var requests []models.Request
	for rows.Next() {
		var req models.Request
		err := rows.Scan(&req.ID, &req.PhoneNumber, &req.FullName, &req.Email, &req.RequestText)
		if err != nil {
			http.Error(w, "Error reading request data", http.StatusInternalServerError)
			return
		}
		requests = append(requests, req)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

// Удаление заявки (доступно разработчикам и администраторам)
func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	type DeleteInput struct {
		RequestID int `json:"request_id"`
	}

	var input DeleteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM requests WHERE id = @p1"
	_, err := database.DB.Exec(query, input.RequestID)
	if err != nil {
		http.Error(w, "Failed to delete request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request deleted successfully"))
}
