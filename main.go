package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nestorsalceda/maverik/devices"
	"github.com/nestorsalceda/maverik/effects"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var ledStrip devices.LedStrip = devices.NewAPA102()

func MyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	effects.Solid(ledStrip, vars["color"])
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/rgb/{color}", MyHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":4567", handlers.CombinedLoggingHandler(os.Stdout, mux)))
}
