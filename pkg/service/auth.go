package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"todo"
	"todo/pkg/repository"
)

const (
	salt       = "skfksdfks12318df"
	signingKey = "aisfbsdbfsfgyeyuabdjhjheiu782q36"
	tokenTTL   = 24 * time.Hour
)

type AuthService struct {
	repo repository.Authorisation
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return claims.UserId, nil
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
