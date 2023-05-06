package main

import (
	"net/http"

	_ "github.com/painhardcore/swag_test/swagger/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", nil)
}
