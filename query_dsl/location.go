package query

type Location struct {
	Type  string     `json:"type"`
	Point [2]float64 `json:"coordinates"` // [lat, long]
}
