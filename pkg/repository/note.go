package repository

import (
	"database/sql"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/jmoiron/sqlx"
)

type PostgresNoteMaker struct {
	db *sqlx.DB
}

func NewPostgresNoteMaker(db *sqlx.DB) *PostgresNoteMaker {
	return &PostgresNoteMaker{db: db}
}

func (np *PostgresNoteMaker) CreateNote(userId int, note butterplanner.Note) (int, error) {
	var noteId int
	var query string
	var row *sql.Row
	if note.Title != "" {
		query = "INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3) RETURNING id"
		row = np.db.QueryRow(query, userId, note.Title, note.Content)
	} else {
		query = "INSERT INTO notes (user_id, content) VALUES ($1, $2) RETURNING id"
		row = np.db.QueryRow(query, userId, note.Content)
	}
	if err := row.Scan(&noteId); err != nil {
		return 0, err
	}
	return noteId, nil
}

func (np *PostgresNoteMaker) GetAllNotes(userId int) ([]butterplanner.Note, error) {
	var notes []butterplanner.Note

	query := "SELECT id, user_id, title, content FROM notes WHERE user_id = $1"
	rows, err := np.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currentNote butterplanner.Note
		if err := rows.Scan(&currentNote.Id, &currentNote.UserId, &currentNote.Title, &currentNote.Content); err != nil {
			return nil, err
		}
		notes = append(notes, currentNote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
