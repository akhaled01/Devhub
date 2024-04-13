package posts

import (
	"encoding/json"
	"net/http"

	"RTF/types"
)

/*
	This is where we get All Posts with No Order
		GET /posts/all
	Refer to the types dir for more info on what a post is
*/
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&types.Post{})
}
