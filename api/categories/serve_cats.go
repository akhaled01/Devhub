package categories

import (
	"RTF/storage/interfaces/categories"
	"RTF/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Serve_categories_handler(w http.ResponseWriter, r *http.Request) {
	catigories, err := categories.GetAllCategoryInfo()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get category information: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(catigories); err != nil {
		utils.ErrorConsoleLog("error encoding all catigories")
		utils.PrintErrorTrace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
