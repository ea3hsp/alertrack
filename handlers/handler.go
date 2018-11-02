package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ea3hsp/alertrack/controllers"

	"github.com/ea3hsp/alertrack/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func tracking(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Message
	msg := `{"message": "Ok message recived. Thank you !"}`
	w.Write([]byte(msg))
	return
}

func lastLocation(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Returns last driver location
	w.Write(location)

	return
}

func homeDistance(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCors(&w)
	// driver query param
	driver := r.URL.Query().Get("driver")
	// create controller
	controller := controllers.NewController()
	// Get driver distance from home
	distance, err := controller.GetDistanceFromHome(driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Error lastlocation getting driver last location : %v", err)
		return
	}
	// Distance container
	var dist = &models.Distance{
		Kilometers: distance / 1000,
	}
	// Marshal distance
	jsonRaw, err := json.Marshal(&dist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Error lastlocation getting driver last location : %v", err)
		return
	}
	// Returns  Status OK (200)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Returns driver distance from home
	w.Write(jsonRaw)
	return
}

// NewHandler ...
func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/driver/tracking", tracking)
	mux.HandleFunc("/api/v1/driver/lastlocation", lastLocation)
	mux.HandleFunc("/api/v1/driver/distance", homeDistance)

	return mux
}
