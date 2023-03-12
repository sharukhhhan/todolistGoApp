package service

import (
	todo "To-do-list"
	"To-do-list/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "sdcfliogtjreoim234#$#"

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
