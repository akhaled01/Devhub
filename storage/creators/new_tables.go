package creators

import (
	"database/sql"
	"errors"
)

var Tables = []func(db *sql.DB) error{
	CreateUserTables,
	CreatePostTables,
	CreateCommentTables,
}

// In the case the database is not found, 
// this function will run to create a new 
// sqlite3 DB file and initialize the
// DB schema of the project.
func CreateNewDB() error {
	DB, err := sql.Open("sqlite3", "RTF.db")
	if err != nil {
		return errors.Join(errors.New("DB creation failed"), err)
	}

	// create tables from the given functions
	for _, table := range Tables {
		if err := table(DB); err != nil {
			return errors.Join(errors.New("table creation failed"), err)
		}
	}

	return nil
}
