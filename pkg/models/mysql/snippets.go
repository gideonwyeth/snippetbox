package mysql

import (
	"database/sql"

	"github.com/gideonwyeth/snippetbox/pkg/models"
)

// define a SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// this will insert a new snippet into the db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	result, err := m.DB.Exec(`INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// this will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}

	err := m.DB.QueryRow(`SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?`, id).Scan(
		&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// this will return last 10 snippets created
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	rows, err := m.DB.Query(`SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
