package mongo

import (
	"encoding/json"
)

type parser struct {
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) ParseRaw(rawLog []byte) (*LogMessage, error) {
	lm := new(LogMessage)
	if err := json.Unmarshal(rawLog, lm); err != nil {
		return nil, err
	}
	return lm, nil
}
