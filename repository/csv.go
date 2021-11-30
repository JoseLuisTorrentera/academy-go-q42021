package repository

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type SpellRepo struct {
	file string
}

func NewSpellRepo(file string) SpellRepo {
	return SpellRepo{file: file}
}

func (sr SpellRepo) GetSpellsFromCsv() ([]*models.Spell, error) {
	csvFile, err := os.Open(sr.file)
	defer csvFile.Close()
	if err != nil {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	spells := []*models.Spell{}
	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		level, err := strconv.Atoi(line[3])
		if err != nil {
			return nil, err
		}

		spell := models.Spell{
			ID:      id,
			Name:    line[1],
			Classes: line[2],
			Level:   level,
			School:  line[4],
		}
		spells = append(spells, &spell)
	}

	return spells, nil
}
