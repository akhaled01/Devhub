package creators

import "database/sql"

func CreateChatTables(DB *sql.DB) error {
	if _, err := DB.Exec(`CREATE TABLE chat_messages ( 
		id                   INTEGER NOT NULL  PRIMARY KEY  ,
		"text"               VARCHAR(250) NOT NULL    ,
		created_at           TIMESTAMP NOT NULL    ,
		sender               VARCHAR(250) NOT NULL    ,
		recipient            VARCHAR(250) NOT NULL    
	 )		
	`); err != nil {
		return err
	}

	return nil
}
