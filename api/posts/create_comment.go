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
		session_id : "xxxxxxxxxxxxx-xxxxxxxxxxx-xxxxx",
		post_id : "xxxxxxxxxxxx-xxxxxxxxxxxxx-xxxxx",
		comment_text : "bla bla bla whatever",
	}

returns 201 on success, 500 on error and 400 on bad request
*/
func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment_creation_request := types.CommentCreationRequest{}

	if err := json.NewDecoder(r.Body).Decode(&comment_creation_request); err != nil {
		utils.ErrorConsoleLog("error decoding json")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := comment.SaveCommentInDB(comment_creation_request); err != nil {
		utils.ErrorConsoleLog("error saving comment")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
