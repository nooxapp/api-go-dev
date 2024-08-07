package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return "", errors.New("JWT_KEY is empty")
	}

	expirationTime := time.Now().Add(time.Minute * 15)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetSession(r *http.Request) (*jwt.RegisteredClaims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, fmt.Errorf("Unauthorized")
		}
		return nil, err
	}

	tokenStr := cookie.Value
	jwtKey := os.Getenv("JWT_KEY")
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
