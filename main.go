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

func RGBHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	effects.Solid(ledStrip, vars["color"])
}

func EnableHandler(w http.ResponseWriter, r *http.Request) {
	effects.Solid(ledStrip, "FFFFFF")
}

func DisableHandler(w http.ResponseWriter, r *http.Request) {
	effects.Solid(ledStrip, "000000")
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/rgb/{color}", RGBHandler).Methods("POST")
	mux.HandleFunc("/enable", EnableHandler).Methods("POST")
	mux.HandleFunc("/disable", DisableHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":4567", handlers.CombinedLoggingHandler(os.Stdout, mux)))
}
