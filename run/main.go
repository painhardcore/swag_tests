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

// home is a handler function
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
