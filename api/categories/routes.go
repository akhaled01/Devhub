package categories

import (
	"net/http"

	"RTF/api/middleware"
)

var Routes = map[string]func(w http.ResponseWriter, r *http.Request){
	"GET /categories":     Serve_categories_handler,
	"OPTIONS /categories": Serve_categories_handler,
}

// Register Post Routes with middleware validation
func RegisterCategoriesRoutes(mux *http.ServeMux) {
	for route, f := range Routes {
		validated := middleware.SessionValidationMiddleware(http.HandlerFunc(f))
		allowed_cors := middleware.AllowCorsMiddleware(validated)
		mux.Handle(route, allowed_cors)
	}
}
