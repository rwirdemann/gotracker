package rest

import (
	"encoding/json"

	"github.com/rwirdemann/go-tracker/design/domain"
)

type JSONConsumer struct {
}

func NewJSONConsumer() JSONConsumer {
	return JSONConsumer{}
}

func (this JSONConsumer) Consume(i interface{}) interface{} {
	var p domain.Project
	json.Unmarshal(i.([]byte), &p)
	return p
}
