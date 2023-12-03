package main

import (
	"net/http"

	"server/handler"
	"server/utils"
)


func main() {
	utils.Ln("By SingSongZepe")
	utils.Ln("Hello, SGetting!")

	// html handler
	http.HandleFunc("/sgetting", handler.SGettingHandler)

	// api handler
	http.HandleFunc("/sgetting/api/ask", handler.AskAPIHandler)

	// main handler
	// e.g. /sgetting/res/img/ai.png
	http.HandleFunc("/", handler.MainHandler)

	http.ListenAndServe(":8080", nil)
}
