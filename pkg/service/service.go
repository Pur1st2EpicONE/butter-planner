package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
)

type ServiceProvider interface {
	CreateUser(user butterplanner.User) (int, error)
}

type Service struct {
	ServiceProvider
}

func NewService(storage *repository.Storage) *Service {
	return &Service{ServiceProvider: NewAuthService(storage)}
}
