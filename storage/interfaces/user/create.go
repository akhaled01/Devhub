package user

import (
	"errors"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
	"github.com/mattn/go-sqlite3"
)

const NEWUSERQUERY = `INSERT INTO users (user_id, user_email, user_name, 
	first_name, last_name, avatar_path, user_pwd) VALUES (?, ?, ?, ?, ?, ?, ?)`

var ErrUserExist = errors.New("a user with either either the username/email already exists")

/*
This function takes in a signup request and creates a
new user in the DB

returns nil if no errors
*/
func SaveUserInDB(u types.SignupRequest) error {
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.Join(errors.New("error Hashing user password"), err)
	}

	stmt, err := storage.DB_Conn.Prepare(NEWUSERQUERY)
	if err != nil {
		return errors.Join(errors.New("error preparing saveUserInDB query"), err)
	}
	defer stmt.Close()

	new_user_id, err := uuid.NewV7()
	if err != nil {
		return errors.Join(errors.New("error creating uuid"), err)
	}

	if _, err := stmt.Exec(new_user_id.String(), u.Email, u.Username, u.FirstName, u.LastName, u.Avatar, hashedPass); err != nil {
		sqliteErr, ok := err.(sqlite3.Error) // extract constraint violation error
		if ok && sqliteErr.Code == sqlite3.ErrConstraint {
			return ErrUserExist
		}
		return errors.Join(errors.New("error executing SaveUserInDB query"), err)
	}

	return nil
}
