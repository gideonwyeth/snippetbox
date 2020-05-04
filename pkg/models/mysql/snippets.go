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
	return 0, nil
}

// this will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// this will return last 10 snippets created
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
