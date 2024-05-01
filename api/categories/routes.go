package categories

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"GET /categories/all": Serve_categories_handler,
}

// Register Post Routes with middleware validation
func RegisterCategoriesRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		mux.Handle(route, middleware.SessionValidationMiddleware(http.HandlerFunc(f)))
	}
}
