package geo

import "math"

// Geo struct
type Geo struct{}

// R metres
const R = 6371e3

// NewGeo creates new geo container
func NewGeo() *Geo {
	return &Geo{}
}

// GetDistance between 2 points
func (g *Geo) GetDistance(pointA [2]float64, pointB [2]float64) float64 {
	// source: https://www.movable-type.co.uk/scripts/latlong.html

	lat1 := pointA[0]
	lat2 := pointB[0]

	lon1 := pointA[1]
	lon2 := pointB[1]

	var φ1 = lat1 * math.Pi / 180
	var φ2 = lat2 * math.Pi / 180

	var Δφ = (lat2 - lat1) * math.Pi / 180
	var Δλ = (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c

}
