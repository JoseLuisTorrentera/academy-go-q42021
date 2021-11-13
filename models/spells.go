package models

type Spell struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Classes string `json:"classes"`
	Level   string `json:"level"`
	School  string `json:"school"`
}
