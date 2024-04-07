package storage

import (
	"database/sql"
	"os"

	"RTF/storage/creators"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB_Conn       *sql.DB
	database_path = "RTF.db"
)

// Initializes connection to database, creates new
// DB file with all the tables if it does not exist, and handles errors gracefully.
func Init() error {
	if _, err := os.Stat(database_path); os.IsNotExist(err) {
		if err := creators.CreateNewDB(); err != nil {
			return err
		}
	}

	db, err := sql.Open("sqlite3", database_path)
	if err != nil {
		return err
	}

	DB_Conn = db
	return nil
}
