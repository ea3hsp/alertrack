package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ea3hsp/alertrack/controllers"

	"github.com/ea3hsp/alertrack/models"
)

func tracking(w http.ResponseWriter, r *http.Request) {
	// Driver location struct definition
	var dl models.DriverLocation
	// Decode incomming http message
	if err := json.NewDecoder(r.Body).Decode(&dl); err != nil {
		log.Printf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	// create controller
	controller := controllers.NewController()
	// Set driver location
	controller.SetDriverLocation(dl)
	// Returns  Status OK (200)
	w.WriteHeader(http.StatusOK)
	return
}

func lastLocation(w http.ResponseWriter, r *http.Request) {
	// driver query param
	driver := r.URL.Query().Get("driver")
	// create controller
	controller := controllers.NewController()
	// Set driver location
	location, err := controller.GetDriverLastLocation(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Error lastlocation getting driver last location : %v", err)
		return
	}
	// Returns  Status OK (200)
	w.WriteHeader(http.StatusOK)
	// Returns last driver location
	w.Write(location)

	return
}

// NewHandler ...
func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tracking", tracking)
	mux.HandleFunc("/lastlocation", lastLocation)

	return mux
}
