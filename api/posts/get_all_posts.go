package posts

import (
	"encoding/json"
	"net/http"

	"RTF/storage/interfaces/post"
	"RTF/utils"
)

/*
This is where we get All Posts with No Order
Responds with a json array

	`GET /posts/all`
*/
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := post.AllPostsFromDB()
	if err != nil {
		utils.ErrorConsoleLog("error getting all posts")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		utils.ErrorConsoleLog("error encoding all posts")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
