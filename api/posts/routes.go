package posts

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"GET /post/all":           GetAllPosts,
	"GET /post/{id}":          GetPostByID,
	"GET /comments/{id}":      GetCommentsByPostID,
	"POST /post/create":       CreatePost,
	"POST /comment/create":    CreateComment,
	"OPTIONS /post/all":       GetAllPosts,
	"OPTIONS /post/create":    CreatePost,
	"OPTIONS /comment/create": CreateComment,
	"OPTIONS /comments/{id}":  GetCommentsByPostID,
}

// Register Post Routes with middleware validation
func RegisterPostRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		validated := middleware.SessionValidationMiddleware(http.HandlerFunc(f))
		allowed_cors := middleware.AllowCorsMiddleware(validated)
		mux.Handle(route, allowed_cors)
	}
}
