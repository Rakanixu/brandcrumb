package fish

// JSON {"id": 7, "tank": 23, "name": "Juan Carlos", "type": "shark", "price": 17.25}
type Fish struct {
	Id       int64   `json:"id"`
	Tank     int64   `json:"tank"`
	Name     string  `json:"name"`
	FishType string  `json:"type"`
	Price    float64 `json:"price"`
}
