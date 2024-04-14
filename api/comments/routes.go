package comments

import "net/http"

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/comments/{id}": GetCommentsByPost,
}

func RegisterCommentRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.HandleFunc(route, f)
	}
}
