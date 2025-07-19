package repository

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/jmoiron/sqlx"
)

type Storer interface {
	CreateUser(user butterplanner.User) (int, error)
	GetUserId(user butterplanner.LoginPassword) (int, error)
}

type NoteMaker interface {
	CreateNote(userId int, note butterplanner.Note) (int, error)
	GetAllNotes(userId int) ([]butterplanner.Note, error)
}

type Storage struct {
	Storer
	NoteMaker
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{Storer: NewPostgresStorer(db),
		NoteMaker: NewPostgresNoteMaker(db)}
}
