package posts

import (
	"encoding/json"
	"net/http"

	"RTF/storage/interfaces/post"
	"RTF/types"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

/*
This is where we get All Posts (newest first)
Responds with a json array

	`GET /posts/all`

	Example successful json response (an array)

	[
		{
			id : "xxxxxxxxxxxxx-xxxxxxxxx-xxxx",
			user : {
				...check types.user json struct
			},
			Image_Path : "INSERT BASE64 STRING ENCODING OF AN IMAGE HERE",
			likes : "20",
			comments : [
				... Multiple comments, check types.comments struct
			],
			Category: {
				id : 4,
				name : "technology"
			},
			Content: "I DO NOT SAY BLUH B-BLUH",
			CreationDate: "2023-10-21"
		},
		{
			id : "xxxxxxxxxxxxx-xxxxxxxxx-xxxx",
			user : {
				...check types.user json struct
			},
			Image_Path : "INSERT BASE64 STRING ENCODING OF AN IMAGE HERE",
			likes : "20",
			comments : [
				... Multiple comments, check types.comments struct
			],
			Category: {
				id : 4,
				name : "technology"
			},
			Content: "I DO NOT SAY BLUH B-BLUH",
			CreationDate: "2023-10-21"
		}
	]
*/
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// get the session id from the cookie
	session_id, err := r.Cookie("session_id")
	if err != nil {
		utils.ErrorConsoleLog(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// get the session from the global sessions map
	requester_session, ok := types.Sessions[uuid.FromStringOrNil(session_id.Value)]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// get all posts from the database
	posts, err := post.AllPostsFromDB(requester_session.User)
	if err != nil {
		utils.ErrorConsoleLog("error getting all posts")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		utils.ErrorConsoleLog("error encoding all posts")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
