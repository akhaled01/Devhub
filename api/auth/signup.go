package auth

import (
	"encoding/json"
	"net/http"

	"RTF/log"
	"RTF/storage/interfaces/user"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for a user to signup on the DevHub.

# The Request URI for this handler

	`POST /auth/signup`

# An example JSON body it accepts

	{
		username: "akhaled",
		email : "akhaledlarus@gmail.com",
		password: "azt@345",
		first_name: "Abdulrahman",
		last_name: "Idrees",
		age: 19,
		gender: "M",
		image: "YOUR_BASE64_ENCODED_IMAGE_DATA"
	}
*/
func Signup(w http.ResponseWriter, r *http.Request) {
	log.InfoConsoleLog("recieved signup request")

	signup_data := types.SignupRequest{}

	if err := json.NewDecoder(r.Body).Decode(&signup_data); err != nil {
		log.ErrorConsoleLog("error decoding json on signup! %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	does_username_exist, err := user.CheckUsernameExist(signup_data.Username)
	if err != nil {
		log.ErrorConsoleLog("error checking username exist")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	does_email_exist, err := user.CheckEmailExist(signup_data.Email)
	if err != nil {
		log.ErrorConsoleLog("error checking email exist")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check if the username or email exist
	if does_username_exist || does_email_exist {
		log.WarnConsoleLog("user tried to signup with existing credentials")
		w.WriteHeader(http.StatusConflict)
		return
	}

	// save avatar image
	img_path, err := utils.SaveImage(signup_data.Avatar, "avatar")
	if err != nil {
		log.ErrorConsoleLog("error saving avatar -> %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signup_data.Avatar = img_path

	if err = user.SaveUserInDB(signup_data); err != nil {
		log.ErrorConsoleLog("error saving user in DB")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
