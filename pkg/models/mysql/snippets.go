package mysql

import (
	"database/sql"

	"github.com/hisyntax/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	//use the LastInsertId() method on the resukt object to get the Id
	//for a newly inserted record in the snippet table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	//the ID return has the type int64, so we convert it
	//to an int type before returning
	return int(id), nil
}

// This will return a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return !0 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
