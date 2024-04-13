package user

import (
	"errors"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

const NEWUSERQUERY = "INSERT INTO users (user_id, user_email, user_name, first_name, last_name, avatar_path, user_pwd) VALUES (?, ?, ?, ?, ?, ?, ?)"

/*
This function takes in a user struct and saves
the instance in the Database.

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
		return errors.Join(errors.New("error executing SaveUserInDB query"), err)
	}

	return nil
}
