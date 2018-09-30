package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ea3hsp/alertrack/database"
	"github.com/ea3hsp/alertrack/models"
)

// Controllers struct
type Controllers struct{}

// NewController returns controller instance
func NewController() *Controllers {
	return &Controllers{}
}

// SetDriverLocation add drive location into database
func (c *Controllers) SetDriverLocation(driverLocation models.DriverLocation) error {
	// Database URL
	db := database.NewDataBase("10.1.2.23", "5984", "drivers")
	// Convert to JSON to store in database
	rawJSON, err := json.Marshal(driverLocation)
	if err != nil {
		log.Println(err)
		return err
	}
	// Create the POST request (insert in database)
	req, err := http.NewRequest("POST", db.URL(), bytes.NewBuffer(rawJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	return err
}

// GetDriverLastLocation get current driver location
func (c *Controllers) GetDriverLastLocation(driver string) ([]byte, error) {

	// Database URL
	db := database.NewDataBase("10.1.2.23", "5984", "drivers")
	view := "/_design/drivers/_view/by_driver"
	url := db.URL() + view

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", driver)
	q.Add("limit", "1")
	q.Add("descending", "true")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return bodyBytes, nil
	}

	return nil, err
}
