package user

import (
	"database/sql"
	"errors"
	"fmt"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"
)

const QUERY_USER = `SELECT * FROM users WHERE %s = ?`

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrPrepare           = errors.New("error preparing statment")
	ErrScan              = errors.New("error scanning results")
)

/*
This Function recieves a field, and a wanted value and interfaces with the database to get
a SINGLE USER by that field.
*/
func GetSingleUser(field, val string) (types.User, error) {
	query := fmt.Sprintf(QUERY_USER, field)
	stmt, err := storage.DB_Conn.Prepare(query)
	if err != nil {
		return (types.User{}), errors.Join(ErrPrepare, err)
	}

	user := types.User{}

	if err := stmt.QueryRow(val).Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Avatar, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return (types.User{}), ErrUserNotFound
		}
		return (types.User{}), errors.Join(ErrScan, err)
	}

	return user, nil
}

/*
This function Authenticates and authorizes access to the user
To be used in the API to authenticate web requests. It accepts username / email
*/
func Authenticate(credential string, password string) (types.User, error) {
	authorized_user := types.User{}

	// Query based on credentials
	var err error
	if utils.IsValidEmail(credential) {
		authorized_user, err = GetSingleUser("user_email", credential)
	} else {
		authorized_user, err = GetSingleUser("user_name", credential)
	}

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return types.User{}, ErrUserNotFound
		}
		return types.User{}, err
	}

	if !utils.CheckPasswordHash(password, authorized_user.Password) {
		return (types.User{}), ErrIncorrectPassword
	}

	return authorized_user, nil
}
