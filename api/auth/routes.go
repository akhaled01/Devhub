package auth

import (
	"net/http"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"POST /auth/signup": Signup,
	"POST /auth/login":  Login,
}

func RegisterAuthRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.HandleFunc(route, f)
	}
}
