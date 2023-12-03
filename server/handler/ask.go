package handler

import (
	"net/http"
	
	"server/utils"
)

// GET /sgetting/api/ask 
func AskAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return 
	}
	utils.Ln("user connected to api ask ...")
	
}