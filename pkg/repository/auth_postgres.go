package repository

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (ps *PostgresStorage) CreateUser(user butterplanner.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, last_name, username, password) VALUES ($1, $2, $3, $4) RETURNING id"

	row := ps.db.QueryRow(query, user.Name, user.Last_name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
