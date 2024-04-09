package creators

import (
	"database/sql"
	"errors"
	"log"
)

var Tables = []func(db *sql.DB) error{
	CreateUserTables,
	CreatePostTables,
	CreateCommentTables,
}

func CreateNewDB() error {
	DB, err := sql.Open("sqlite3", "RTF.db")
	if err != nil {
		log.Fatal(err)
	}
	
	// create tables from the given functions
	for _, table := range Tables {
		if err := table(DB); err != nil {
			return errors.Join(errors.New("table creation failed"), err)
		}
	}

	return nil
}
