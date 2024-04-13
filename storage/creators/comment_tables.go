package creators

import (
	"database/sql"
)

func CreateCommentTables(DB *sql.DB) error {
	TableStatements := []string{
		`CREATE TABLE IF NOT EXISTS comments (
            comm_id VARCHAR NOT NULL PRIMARY KEY,
            post_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            comment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            comment VARCHAR(150) NOT NULL,
            edited BOOLEAN NOT NULL DEFAULT FALSE,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (user_id) REFERENCES users(user_id)
        )`,
		`CREATE TABLE IF NOT EXISTS comment_likes (
            action_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            comment_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            actions_type BOOLEAN NOT NULL,
            action_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (comment_id) REFERENCES comments(comm_id),
            FOREIGN KEY (user_id) REFERENCES users(user_id)
        )`,
	}

	for _, stmt := range TableStatements {
		if _, err := DB.Exec(stmt); err != nil {
			return err
		}
	}

	return nil
}
