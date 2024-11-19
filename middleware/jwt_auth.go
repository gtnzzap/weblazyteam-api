package middleware

import (
	"context"
	"net/http"
	"strings"
	"weblazyteam-api/utils"
)

// Ключи для передачи данных контекста
type contextKey string

const (
	UserLoginContextKey contextKey = "userLogin"
	UserRoleContextKey  contextKey = "userRole"
)

// Middleware для проверки JWT токена
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получение токена из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header missing or malformed", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Проверка и декодирование токена
		claims, err := utils.ParseJWT(token)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Передача данных из токена в контекст запроса
		ctx := context.WithValue(r.Context(), UserLoginContextKey, claims.Login)
		ctx = context.WithValue(ctx, UserRoleContextKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Middleware для проверки роли пользователя
func RoleAuth(requiredRole string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(UserRoleContextKey).(string)
		if role != requiredRole && role != "admin" {
			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
