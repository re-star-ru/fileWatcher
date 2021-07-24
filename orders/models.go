package orders

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type Order struct {
	id    string
	Name  string `json:"name"`
	Stand string `json:"stand"` //

	Car  string `json:"car"`
	Unit string `json:"unit"` // генератор или стартер
}

func (o *Order) FromJson(r io.Reader) error {
	if err := json.NewDecoder(r).Decode(o); err != nil {
		return err
	}
	if o.Stand == "" {
		o.Stand = "ASG_113.local" // default stand
	}

	return o.Validate()
}

func (o Order) Validate() error {
	if o.Name == "" {
		return errors.New("name is empty")
	}
	if o.Stand == "" {
		return errors.New("stand name is empty")
	}

	return nil
}
