package creators

import (
	"database/sql"
)

func CreatePostTables(DB *sql.DB) error {
	TableStatements := []string{
		`CREATE TABLE IF NOT EXISTS posts (
            post_id               VARCHAR NOT NULL PRIMARY KEY,
            user_id               VARCHAR NOT NULL,
            creation_date         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            post_content          VARCHAR(150) NOT NULL,
            post_image_path       VARCHAR(150),
            FOREIGN KEY (user_id) REFERENCES users(user_id)
        )`,
		`CREATE TABLE IF NOT EXISTS category (
            cat_id               INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            category             CHAR NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS post_categories (
            ID                    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            post_id               VARCHAR NOT NULL,
            cat_id                INTEGER NOT NULL DEFAULT 1,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (cat_id)  REFERENCES category(cat_id)
        )`,
		`CREATE TABLE IF NOT EXISTS post_likes (
            like_id                INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            post_id                VARCHAR NOT NULL,
            user_id                VARCHAR NOT NULL,
            FOREIGN KEY (post_id)  REFERENCES posts(post_id),
            FOREIGN KEY (user_id)  REFERENCES users(user_id)
        )`,
		`CREATE TABLE sessions ( 
						id                           INTEGER      NOT NULL PRIMARY KEY AUTOINCREMENT,
						session_key                  VARCHAR(250) NOT NULL,
						created_at                   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
						expiry_time                  TIMESTAMP    NOT NULL,
						user_id                      VARCHAR(255) NOT NULL,
						FOREIGN KEY ( user_id ) REFERENCES users( user_id )  
				)`,
		`CREATE TABLE logs ( 
						id                                INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
						user_id              		          VARCHAR(255) NOT NULL,
						event_id                          VARCHAR(250) NOT NULL,
						event_type                        VARCHAR(250),
						FOREIGN KEY ( user_id ) REFERENCES users( user_id )  
				)`,
	}

	for _, stmt := range TableStatements {
		if _, err := DB.Exec(stmt); err != nil {
			return err
		}
	}

	stmt, err := DB.Prepare("INSERT INTO category (category) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	categories := []string{"General", "Engineering", "Travel", "Technology", "Mathematics"}
	for _, category := range categories {
		if _, err := stmt.Exec(category); err != nil {
			return err
		}
	}
	
	return nil
}
