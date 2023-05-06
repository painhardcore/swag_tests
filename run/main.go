package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

// simple http server in go with handlers
func main() {
	// register the handlers
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	// start the server
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		panic(err)
	}
}

// home is a home function
// @Summary      Home
// @Description  Home page
// @Tags         index
// @Success      200  "Hello from home page"
// @Failure 400 {string} string "error: request is not valid"
// @Failure 404 {string} string "error: no homa data found"
// @Failure 500 {string} string "error: failed to get something "
// @Router       / [get]
func home(w http.ResponseWriter, r *http.Request) {
	log.Info("Home page accessed")
	w.Write([]byte("Hello from home page"))
}

// about is a handler function
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from about page"))
}

// contact is a handler function
func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from contact page"))
}
