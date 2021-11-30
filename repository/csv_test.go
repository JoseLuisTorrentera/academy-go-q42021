package repository

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/stretchr/testify/assert"
)

var spells = []*models.Spell{
	{
		ID:      1,
		Name:    "test spell 1",
		Classes: "test_classes_1",
		Level:   0,
		School:  "test_school_1",
	},
	{
		ID:      2,
		Name:    "test spell 2",
		Classes: "test_classes_2",
		Level:   0,
		School:  "test_school_1",
	},
}

func TestSpellRepo_GetSpellsFromCsv(t *testing.T) {
	createTestFile()

	test_cases := []struct {
		name           string
		expectedLength int
		response       []*models.Spell
		hasError       bool
		errorIs        error
		file           string
	}{
		{
			"get all spells",
			2,
			spells,
			false,
			nil,
			"test.csv",
		},
		{
			"error in repository",
			0,
			nil,
			true,
			errors.New("No file"),
			"",
		},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			sr := NewSpellRepo(tc.file)
			spells, err := sr.GetSpellsFromCsv()

			assert.EqualValues(t, tc.expectedLength, len(spells))
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}

}

func createTestFile() {

	spellsFileData := [][]string{
		{
			"1",
			"test spell 1",
			"test_classes_1",
			"0",
			"test_school_1",
		},
		{
			"2",
			"test spell 2",
			"test_classes_2",
			"0",
			"test_school_1",
		},
	}

	csvFile, err := os.Create("test.csv")
	defer csvFile.Close()
	if err != nil {
		log.Fatalln("Failed creating test csv file", err)
	}

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	w.WriteAll(spellsFileData)

}
