package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

func GetAllSpells() ([]*models.Spell, error) {
	csvFile, err := os.Open("./commons/dnd-spells.csv")
	if err != nil {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	spells := []*models.Spell{}
	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])
		spell := models.Spell{
			Id:      id,
			Name:    line[1],
			Classes: line[2],
			Level:   line[3],
			School:  line[4],
		}
		spells = append(spells, &spell)
	}

	return spells, nil
}

func GetSpellById(spell_id int) (*models.Spell, error) {
	csvFile, err := os.Open("./commons/dnd-spells.csv")
	if err != nil {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])
		if spell_id == id {
			spell := models.Spell{
				Id:      id,
				Name:    line[1],
				Classes: line[2],
				Level:   line[3],
				School:  line[4],
			}
			return &spell, nil
		}
	}

	return nil, errors.New("Spell not founded!")
}
