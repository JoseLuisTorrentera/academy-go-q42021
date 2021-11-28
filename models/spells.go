package models

//Spell - Spell structure
type Spell struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Classes string `json:"classes"`
	Level   int    `json:"level"`
	School  string `json:"school"`
}
