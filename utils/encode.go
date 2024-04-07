package utils

import (
	"encoding/json"
	"net/http"
)

// encodes anything to JSON and writes the response
func EncodeJSON(w http.ResponseWriter, v any) error {
	return json.NewEncoder(w).Encode(v)
}
