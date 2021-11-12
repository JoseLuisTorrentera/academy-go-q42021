package models

type Spell struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	School string `json:"school"`
}
