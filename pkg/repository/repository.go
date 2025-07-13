package repository

import "github.com/jmoiron/sqlx"

type ReposAuthorizer interface {
}

type Repository struct {
	ReposAuthorizer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
