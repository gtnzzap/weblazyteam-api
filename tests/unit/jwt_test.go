package unit

import (
	"testing"
	"weblazyteam-api/utils"
)

func TestGenerateJWT(t *testing.T) {
	claims := &utils.Claims{
		Login: "testuser",
		Role:  "user",
	}
	token, err := utils.GenerateJWT(claims.Login, claims.Role)
	if err != nil {
		t.Fatalf("Ошибка генерации JWT: %v", err)
	}
	if token == "" {
		t.Fatalf("Токен не должен быть пустым")
	}
}

func TestParseJWT(t *testing.T) {
	claims := &utils.Claims{
		Login: "testuser",
		Role:  "user",
	}
	token, _ := utils.GenerateJWT(claims.Login, claims.Role)

	parsedClaims, err := utils.ParseJWT(token)
	if err != nil {
		t.Fatalf("Ошибка парсинга JWT: %v", err)
	}
	if parsedClaims.Login != claims.Login {
		t.Fatalf("Имена пользователей не совпадают")
	}
}
