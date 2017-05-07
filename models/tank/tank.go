package tank

import (
	"github.com/Rakanixu/brandcrumb/models/fish"
)

// JSON {"id": 1, "name":"Mediterranean sea", "volume": 350.00, "temperature": 20.00}
type Tank struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Volume      float64      `json:"volume"`
	Temperature float32      `json:"temperature"`
	Fish        []*fish.Fish `json:"fish,omitempty"` // This is just a helper to generate JSON response. They are not stored together.
}
