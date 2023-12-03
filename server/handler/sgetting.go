package handler

import (
	"net/http"
	
	"server/utils"
	"server/value"
)

func SGettingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusBadRequest)
		return
	}
	utils.Ln("user connected to sgetting")
	
	file, err := utils.ReadFile(value.SGETTING_PATH)
	if err != nil {
		http.Error(w, "Read sgetting file failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(file)
}