package routes

import (
	"net/http"

	"github.com/azizsama/go-web-mvc-starter/controllers"
)

func Start(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Create your routes here...
	mux.HandleFunc("/", controllers.IndexHandler)
}
