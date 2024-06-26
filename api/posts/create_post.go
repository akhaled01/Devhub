package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"RTF/storage/interfaces/post"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for post creation

	`POST /post/create`

json body it accepts

	{
		post_text : "bla bla bla whatever",
		post_image : "INSERT BASE64 ENCODING OF AN IMAGE HERE"
	}

returns 201 on success, 500 on error and 400 on bad request
*/
func CreatePost(w http.ResponseWriter, r *http.Request) {
	utils.InfoConsoleLog("recieved post create Request")
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

	// decode post text from URI
	decodedPostText, err := url.QueryUnescape(post_creation_request.Post_text)
	if err != nil {
		http.Error(w, "Failed to decode post text", http.StatusBadRequest)
		return
	}
	post_creation_request.Post_text = decodedPostText

	post_creation_request.Post_text = strings.ReplaceAll(post_creation_request.Post_text, "\n", "@")

	fmt.Println(post_creation_request.Post_text)

	new_post_object, err := post.ConstructNewPostFromRequest(post_creation_request)
	if err != nil {
		utils.ErrorConsoleLog("error constructing a new post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := post.SavePostInDB(new_post_object, post_creation_request.Post_category); err != nil {
		utils.ErrorConsoleLog("error saving the new post")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
