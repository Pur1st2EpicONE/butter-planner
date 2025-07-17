package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
)

type ServiceProvider interface {
	CreateUser(user butterplanner.User) (int, error)
	GetUserId(user butterplanner.LoginPassword) (int, error)
	CreateToken(userId int) (string, error)
}

type Service struct {
	ServiceProvider
}

func NewService(storage *repository.Storage) *Service {
	return &Service{ServiceProvider: NewAuthService(storage)}
}
