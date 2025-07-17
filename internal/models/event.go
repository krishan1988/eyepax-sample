package models

import "encoding/json"

type Event struct {
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
}

type Events []Event

func (es *Events) Read(data string) error {
	return json.Unmarshal([]byte(data), es)
}
