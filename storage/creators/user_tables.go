package creators

import "database/sql"

func CreateUserTables(DB *sql.DB) error {
	if _, err := DB.Exec(`CREATE TABLE users (
			user_id                 VARCHAR(255) NOT NULL PRIMARY KEY,
			user_email              VARCHAR(50)  NOT NULL,
			user_name               VARCHAR(50)  NOT NULL,
			first_name              VARCHAR(50),
			last_name               VARCHAR(50),
			gender                  VARCHAR(10),
			avatar_path             VARCHAR(255),
			user_pwd                VARCHAR NOT NULL,
			CONSTRAINT unq_emails   UNIQUE ( user_email ),
			CONSTRAINT unq_username UNIQUE ( user_name )
		 )`); err != nil {
		return err
	}

	return nil
}
