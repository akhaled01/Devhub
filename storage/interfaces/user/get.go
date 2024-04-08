package user

import (
	"errors"

	"RTF/storage"
	"RTF/types"

	"github.com/mattn/go-sqlite3"
)

/*
This Function recieves a field, and a wanted value and interfaces with the database to get
a SINGLE USER by that field. A custom error will be returned for effective logging.

Params:

	field string
	val   string
*/
func GetSingleUser(field, val string) (*types.User, error) {
	user := &types.User{}
	stmt, err := storage.DB_Conn.Prepare("SELECT * FROM users WHERE ? = ?")
	if err != nil {
		return nil, errors.Join(errors.New("ERROR PREPARE SELECT STATEMENT"), err)
	}

	if err := stmt.QueryRow(field, val).Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.Avatar, &user.LastName, &user.Password); err != nil {
		if err == sqlite3.ErrNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.Join(errors.New("ERROR SCANNING SELECT RESULT"), err)
	}

	return user, nil
}
