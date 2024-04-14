package chat

import "net/http"

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"/ws": ChatHandler,
}

func RegisterChatRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.HandleFunc(route, f)
	}
}
