package posts

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"GET /post/all":     GetAllPosts,
	"GET /post/{id}":    GetPostByID,
	"POST /post/create": CreatePost,
}

// Register Post Routes with middleware validation
func RegisterPostRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.Handle(route, middleware.SessionValidationMiddleware(http.HandlerFunc(f)))
	}
}
