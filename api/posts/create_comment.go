package posts

import (
	"encoding/json"
	"net/http"

	"RTF/storage/interfaces/comment"
	"RTF/types"
	"RTF/utils"
)

/*
This is the handler for comment creation

	`POST /comment/create`

json body it accepts

	{
		post_id : "xxxxxxxxxxxx-xxxxxxxxxxxxx-xxxxx",
		comment_text : "bla bla bla whatever",
	}

returns 201 on success, 500 on error and 400 on bad request
*/
func CreateComment(w http.ResponseWriter, r *http.Request) {
	utils.InfoConsoleLog("recieved comment creation request")
	comment_creation_request := types.CommentCreationRequest{}

	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Saving the session id in the request (idk why!?)
	comment_creation_request.Session_id = session_id.Value

	if err := json.NewDecoder(r.Body).Decode(&comment_creation_request); err != nil {
		utils.ErrorConsoleLog("error decoding json")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment_obj, err := comment.ConstructNewCommentFromRequest(comment_creation_request)
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		return
	}

	if err := comment.SaveCommentInDB(comment_obj); err != nil {
		utils.ErrorConsoleLog("error saving comment")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
