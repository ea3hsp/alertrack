package models

// Vehicle struct
type Vehicle struct {
	Type       string `json:"type"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Color      string `json:"color"`
	RegisPlate string `json:"regisPlate"`
}

// Driver struct
type Driver struct {
	Name string `json:"name"`
}

// DriverLocation driver location struct
type DriverLocation struct {
	Driver     Driver      `json:"driver"`
	Vehicle    Vehicle     `json:"vehicle"`
	Coordinate [][]float64 `json:"coordinate"`
	Timestamp  int64       `json:"timestamp"`
}
