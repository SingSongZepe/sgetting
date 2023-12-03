package handler

import (
	"net/http"

	"server/value"
)

// main handler
func MainHandler(w http.ResponseWriter, r *http.Request) {
	
	fs := http.FileServer(http.Dir(value.WEB_ROOT))

	fh := http.StripPrefix("/sgetting/", fs)
	
	// return the value
	fh.ServeHTTP(w, r)
	
}