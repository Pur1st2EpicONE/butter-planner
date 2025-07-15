package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
)

type AuthService struct {
	storer repository.Storer
}

func NewAuthService(storage repository.Storer) *AuthService {
	return &AuthService{storer: storage}
}

func (s *AuthService) CreateUser(user butterplanner.User) (int, error) {
	return s.storer.CreateUser(user)
}
