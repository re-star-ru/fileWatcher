package orders

import (
	"encoding/json"
	"io"
)

type Order struct {
	id    string
	Name  string `json:"name"`
	Stand string `json:"stand"`

	Car  string `json:"car"`
	Unit string `json:"unit"` // генератор или стартер
}

func (o *Order) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(o)
}
