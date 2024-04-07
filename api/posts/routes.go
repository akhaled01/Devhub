package posts

import "net/http"

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/post/all": GetAllPosts,
}

func RegisterPostRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.HandleFunc(route, f)
	}
}
