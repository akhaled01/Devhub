package posts

import (
	"encoding/json"
	"net/http"

	"RTF/storage/interfaces/post"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for post creation

	`POST /post/create`

json body it accepts

	{
		session_id : "xxxxxxxxxxxxx-xxxxxxxxxxx-xxxxx",
		post_text : "bla bla bla whatever",
		post_image : "INSERT BASE64 ENCODING OF AN IMAGE HERE"
	}

returns 201 on success, 500 on error and 400 on bad request
*/
func CreatePost(w http.ResponseWriter, r *http.Request) {
	post_creation_request := types.PostCreationRequest{}
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	post_creation_request.Session_id = session_id.Value
	if err := json.NewDecoder(r.Body).Decode(&post_creation_request); err != nil {
		utils.ErrorConsoleLog("error decoding json")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	new_post_object, err := post.ConstructNewPostFromRequest(post_creation_request)

	if err != nil {
		utils.ErrorConsoleLog("error constructing a new post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := post.SavePostInDB(new_post_object); err != nil {
		utils.ErrorConsoleLog("error saving the new post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
