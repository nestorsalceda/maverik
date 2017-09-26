package http

import (
	"net/http"

	"github.com/nestorsalceda/maverik/devices"
	"github.com/nestorsalceda/maverik/effects"

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

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/rgb/{color}", RGBHandler).Methods("POST")
	router.HandleFunc("/enable", EnableHandler).Methods("POST")
	router.HandleFunc("/disable", DisableHandler).Methods("POST")

	return router
}
