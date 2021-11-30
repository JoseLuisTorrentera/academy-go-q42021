package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type UpdateCSV struct {
}

func NewUpdateCSV() UpdateCSV {
	return UpdateCSV{}
}

func (upcsv UpdateCSV) UpdateSpellsCSV(file string, spell *models.Spell) error {
	csvFile, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer csvFile.Close()
	if err != nil {
		return err
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	data := []string{
		strconv.Itoa(spell.ID),
		spell.Name,
		spell.Classes,
		strconv.Itoa(spell.Level),
		spell.School,
	}

	err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
