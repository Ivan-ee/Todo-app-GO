package service

import (
	"crypto/sha1"
	"fmt"
	"todo"
	"todo/pkg/repository"
)

const salt = "skfksdfks12318df"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
