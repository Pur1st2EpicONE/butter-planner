package repository

import (
	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

func (ps *PostgresStorage) GetUserId(user butterplanner.LoginPassword) (int, error) {
	var id int
	var passHash string

	query := "SELECT id, password FROM users WHERE username=$1"
	row := ps.db.QueryRow(query, user.Username)
	if err := row.Scan(&id, &passHash); err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(user.Password)); err != nil {
		return 0, err
	}

	return id, nil
}
