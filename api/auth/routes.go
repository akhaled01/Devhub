package auth

import (
	"RTF/api/middleware"
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"POST /auth/signup":    Signup,
	"POST /auth/login":     Login,
	"OPTIONS /auth/login":  Login,
	"OPTIONS /auth/signup": Signup,
}

func RegisterAuthRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		allowed_cors := middleware.AllowCorsMiddleware(http.HandlerFunc(f))
		mux.Handle(route, allowed_cors)
	}
}
