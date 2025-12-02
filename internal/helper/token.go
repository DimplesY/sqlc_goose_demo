package helper

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	t *jwt.Token
)

var key = []byte("test")

func GenerateToken(name string) (string, error) {
	claims := jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")

		slog.Info("获取登录凭证", "tokenStr", tokenStr)

		if tokenStr == "" {
			slog.Error("未登录")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := VerifyToken(tokenStr)
		if err != nil {
			slog.Error("登录凭证无效", "error", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			slog.Error("登录凭证无效")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		slog.Info("登录凭证有效", "claims", token.Claims)
		next.ServeHTTP(w, r)
	})
}
