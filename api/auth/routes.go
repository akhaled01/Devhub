package auth

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"POST /auth/signup":    Signup,
	"POST /auth/login":     Login,
	"POST /auth/logout":    Logout,
	"OPTIONS /auth/login":  Login,
	"OPTIONS /auth/signup": Signup,
	"OPTIONS /auth/logout": Logout,
}

func RegisterAuthRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		allowed_cors := middleware.AllowCorsMiddleware(http.HandlerFunc(f))
		mux.Handle(route, allowed_cors)
	}
}
