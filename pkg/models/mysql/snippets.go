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
	//write the SQL statement we want to execute
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	row := m.DB.QueryRow(stmt, id)

	//initialize a pointer to a new zeroed snippet struct
	s := &models.Snippet{}

	//use row.Scan() method to copy the values from each field in sql.Row to the
	//corresponding field in the snippe struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	//if everythin went well then return the snippet object
	return s, nil
}

// This will return !0 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	//write the SQL statement we want to execute
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10"

	//use the QUERY() method on the connection pool to execute our SQL statement
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	//initialize an empty slice to hold the models.Snipets objects
	snippets := []*models.Snippet{}

	//use rows.Next to iterate through the rows in the resultset
	for rows.Next() {
		//create a pointer to a new zeroed Snippet struct
		s := &models.Snippet{}

		//use rows.Scan() to copy the values from each field in the rows to the new snippet object that we created
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		//append it to the slice of snippets
		snippets = append(snippets, s)

	}

	//when the rows.Nest() loop has finished, we call rows.Err() to retrieve an error that
	//was ecountered during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//if everything went well, then return the snippets slice
	return snippets, nil

}
