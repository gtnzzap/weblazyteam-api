package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Секрет для подписи токенов
var jwtSecret = []byte("your_secret_key")

// Кастомные claims для JWT
type Claims struct {
	Login string `json:"login"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// Генерация JWT токена
func GenerateJWT(login, role string) (string, error) {
	claims := Claims{
		Login: login,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Декодирование и проверка JWT токена
func ParseJWT(tokenString string) (*Claims, error) {
	// Парсим токен и получаем ошибку, если она возникла
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// возвращаем секретный ключ для верификации токена
		return jwtSecret, nil
	})

	// Проверка на ошибку парсинга токена
	if err != nil {
		return nil, err // Возвращаем ошибку если парсинг не удался
	}

	// Проверка на корректность токена и извлечение данных
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
