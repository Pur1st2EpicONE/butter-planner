package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	storer repository.Storer
}

func NewAuthService(storer repository.Storer) *AuthService {
	return &AuthService{storer: storer}
}

func (s *AuthService) CreateUser(user butterplanner.User) (int, error) {
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	return s.storer.CreateUser(user)
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}
