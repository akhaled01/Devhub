package likes

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"POST /likePost/{id}":       LikePost,
	"OPTIONS /likePost/{id}":    LikePost,
	"GET /userstats":            GetUserCounts,
	"POST /likeComment/{id}":    LikeComment,
	"OPTIONS /likeComment/{id}": LikeComment,
	"OPTIONS /userstats":        GetUserCounts,
}

// Register Post Routes with middleware validation
func RegisterLikeRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		validated := middleware.SessionValidationMiddleware(http.HandlerFunc(f))
		allowed_cors := middleware.AllowCorsMiddleware(validated)
		mux.Handle(route, allowed_cors)
	}
}
