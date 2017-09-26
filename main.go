package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	maverik_http "github.com/nestorsalceda/maverik/http"
)

func main() {
	router := maverik_http.NewRouter()

	log.Fatal(
		http.ListenAndServe(":4567",
			handlers.CombinedLoggingHandler(os.Stdout, router)))
}
