package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword принимает строку пароля и возвращает его хешированную версию.
// Возвращает хеш и возможную ошибку.
func HashPassword(password string) (string, error) {
	// Используем bcrypt для хеширования пароля с заданной стоимостью (по умолчанию: 10).
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Возвращаем хеш в виде строки.
	return string(bytes), nil
}

// CheckPasswordHash сравнивает исходный пароль с его хешем.
// Возвращает true, если пароли совпадают, иначе false.
func CheckPasswordHash(password, hash string) bool {
	// bcrypt.CompareHashAndPassword возвращает nil, если пароль соответствует хешу.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
