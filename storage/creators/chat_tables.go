package creators

import "database/sql"

func CreateChatTables(DB *sql.DB) error {
	if _, err := DB.Exec(`CREATE TABLE chats ( 
		id                   VARCHAR(250) NOT NULL  PRIMARY KEY  ,
		client_id            INTEGER NOT NULL    ,
		created_at           TIMESTAMP     ,
		FOREIGN KEY ( client_id ) REFERENCES chat_clients( id )  
	 );
	`); err != nil {
		return err
	}

	if _, err := DB.Exec(`CREATE TABLE chat_messages ( 
		id                   INTEGER NOT NULL  PRIMARY KEY  ,
		chat_id              VARCHAR(250) NOT NULL    ,
		client_id            VARCHAR(250) NOT NULL    ,
		"text"               VARCHAR(250)     ,
		created_at           TIMESTAMP     ,
		FOREIGN KEY ( chat_id ) REFERENCES chats( id )  ,
		FOREIGN KEY ( client_id ) REFERENCES chat_clients( id )  
	 );	
	`); err != nil {
		return err
	}

	if _, err := DB.Exec(`CREATE TABLE chat_clients ( 
		id                   INTEGER NOT NULL  PRIMARY KEY  ,
		user_id              VARCHAR(250) NOT NULL    
	 );	
	`); err != nil {
		return err
	}

	return nil
}
