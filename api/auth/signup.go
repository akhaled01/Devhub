package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"RTF/log"
	"RTF/storage/interfaces/user"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for a user to signup on the DevHub.

# The Request URI for this handler (returns 201 Created is successful)

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
		log.ErrorConsoleLog("error decoding json on signup!")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// save avatar image
	img_path, err := utils.SaveImage(signup_data.Avatar, "avatar")
	if err != nil {
		log.ErrorConsoleLog("error saving avatar")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signup_data.Avatar = img_path

	if err = user.SaveUserInDB(signup_data); err != nil {
		if errors.Is(err, user.ErrUserExist) {
			log.WarnConsoleLog("a user with either this username / email exist")
			w.WriteHeader(http.StatusConflict)
			return
		}
		log.ErrorConsoleLog("error saving user in DB")
		log.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
