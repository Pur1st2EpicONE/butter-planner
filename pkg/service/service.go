package service

import "github.com/Pur1st2EpicONE/butter-planner/pkg/repository"

type ServiceAuthorizer interface {
}

type Service struct {
	ServiceAuthorizer
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
