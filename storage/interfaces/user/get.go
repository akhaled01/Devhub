package user

import (
	"database/sql"
	"errors"

	"RTF/storage"
	"RTF/types"
	"RTF/utils"
)

const QUERY_USER = `SELECT * FROM users WHERE ? = ?`

/*
This Function recieves a field, and a wanted value and interfaces with the database to get
a SINGLE USER by that field. A custom error will be returned for effective logging.

Params:

	field string
	val   string
*/
func GetSingleUser(field, val string) (*types.User, error) {
	user := &types.User{}
	stmt, err := storage.DB_Conn.Prepare(QUERY_USER)
	if err != nil {
		return nil, errors.Join(errors.New("ERROR PREPARING SELECT STATEMENT"), err)
	}

	if err := stmt.QueryRow(field, val).Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.Avatar, &user.LastName, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Join(errors.New("ERROR SCANNING SELECT RESULT"), err)
	}

	return user, nil
}

/*
This function Authenticates and authorizes access to the user
To be used in the API to authenticate web requests. It accepts username / email

Regex is used to parse the credential and check if its a email, and query the DB based on
that
*/
func Authenticate(credential string, password string) (*types.User, error) {
	authorized_user := &types.User{}

	// query based on credentials
	if utils.IsValidEmail(credential) {
		user, err := GetSingleUser("user_email", credential)
		if err != nil {
			if errors.Is(err, errors.New("user not found")) {
				return nil, errors.New("user not found")
			}
		}
		authorized_user = user
	} else {
		user, err := GetSingleUser("user_name", credential)
		if err != nil {
			if errors.Is(err, errors.New("user not found")) {
				return nil, errors.New("user not found")
			}
		}
		authorized_user = user
	}

	h, err := utils.HashPassword(authorized_user.Password)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, h) {
		return nil, errors.New("password incorrect, unauthorized access")
	}

	return authorized_user, nil
}

func CheckUsernameExist(username string) (bool, error) {
	user, _ := GetSingleUser("user_name", username)

	return user != nil, nil
}

func CheckEmailExist(email string) (bool, error) {
	user, _ := GetSingleUser("user_email", email)

	return user != nil, nil
}
