package service

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
)

type ServiceProvider interface {
	CreateUser(user butterplanner.User) (int, error)
	GetUserId(user butterplanner.LoginPassword) (int, error)
	CreateToken(userId int) (string, error)
	ParseToken(token string) (int, error)
}

type NoteServiceProvider interface {
	CreateNote(userId int, note butterplanner.Note) (int, error)
	GetAllNotes(userId int) ([]butterplanner.Note, error)
}

type Service struct {
	ServiceProvider
	NoteServiceProvider
}

func NewService(storage *repository.Storage) *Service {
	return &Service{ServiceProvider: NewAuthService(storage),
		NoteServiceProvider: NewNoteService(storage.NoteMaker)}
}
