package repository

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/jmoiron/sqlx"
)

type Storer interface {
	CreateUser(user butterplanner.User) (int, error)
	GetUserId(user butterplanner.LoginPassword) (int, error)
}

type Storage struct {
	Storer
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{Storer: NewPostgresStorage(db)}
}
