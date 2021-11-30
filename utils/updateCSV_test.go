package utils

import (
	"errors"
	"testing"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/stretchr/testify/assert"
)

var spell models.Spell = models.Spell{
	ID:      1,
	Name:    "test spell 1",
	Classes: "test_classes_1",
	Level:   0,
	School:  "test_school_1",
}

func TestUpdateCSV_UpdateSpellsCSV(t *testing.T) {
	testCases := []struct {
		name     string
		hasError bool
		errorIs  error
		file     string
	}{
		{
			name:     "update succesful",
			hasError: false,
			errorIs:  nil,
			file:     "test.csv",
		},
		{
			name:     "update failure; error file",
			hasError: true,
			errorIs:  errors.New("Error updating csv"),
			file:     "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			upCsv := NewUpdateCSV()
			err := upCsv.UpdateSpellsCSV(tc.file, &spell)

			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
