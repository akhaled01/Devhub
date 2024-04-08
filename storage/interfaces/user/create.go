package user

import (
	"errors"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"
)

const NEWUSERQUERY = "INSERT INTO users VALUE (user_email, user_name, first_name, last_name, avatar_path, user_pwd) VALUES (?, ?, ?, ?, ?, ?)"

/*
This function takes in a user struct and saves
the instance in the Database.

returns nil if no errors
*/
func SaveUserInDB(u types.User) error {
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.Join(errors.New("ERROR Hashing user password"), err)
	}

	stmt, err := storage.DB_Conn.Prepare(NEWUSERQUERY)
	if err != nil {
		return errors.Join(errors.New("ERROR preparing saveUserInDB query"), err)
	}

	if _, err := stmt.Exec(u.Email, u.Username, u.FirstName, u.LastName, u.Avatar, hashedPass); err != nil {
		return errors.Join(errors.New("ERROR executing SaveUserInDB QUERY"), err)
	}

	return nil
}
