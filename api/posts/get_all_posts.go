package posts

import (
	"encoding/json"
	"net/http"

	"RTF/storage/interfaces/post"
	"RTF/utils"
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
	posts, err := post.AllPostsFromDB()
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
