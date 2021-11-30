package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var spell = models.Spell{
	ID:      1,
	Name:    "test spell",
	Classes: "test_classes,",
	Level:   1,
	School:  "test_school",
}

var apiResponse string = `{"name":"test spell","level":1,"school":{"name":"test_school"},"classes":[{"name":"test_classes"}]}`

func TestSpellApiService_GetSpellByName(t *testing.T) {
	createTestFile()

	testCases := []struct {
		name        string
		response    *models.Spell
		hasError    bool
		errorIs     error
		spellName   string
		apiResponse string
		apiCode     int
		file        string
	}{
		{
			name:        "successful request",
			response:    &spell,
			hasError:    false,
			errorIs:     nil,
			spellName:   "test-spell",
			apiResponse: apiResponse,
			apiCode:     200,
			file:        "test.csv",
		},
		{
			name:        "failed request",
			response:    nil,
			hasError:    true,
			errorIs:     errors.New("No spell by that name"),
			spellName:   "non existing spell",
			apiResponse: "",
			apiCode:     404,
			file:        "test.csv",
		},
		{
			name:        "failed id generation",
			response:    nil,
			hasError:    true,
			errorIs:     errors.New("No existing file"),
			spellName:   "test-spell",
			apiResponse: apiResponse,
			apiCode:     200,
			file:        "",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("https://www.dnd5eapi.co/api/spells/%s", tc.spellName)
			service := NewSpellApiService(tc.file)
			httpmock.ActivateNonDefault(service.Client.GetClient())
			defer httpmock.DeactivateAndReset()

			responder := httpmock.NewBytesResponder(tc.apiCode, []byte(tc.apiResponse))
			httpmock.RegisterResponder("GET", url, responder)

			s, err := service.GetSpellByName(tc.spellName)
			assert.EqualValues(t, tc.response, s)
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
			"0",
			"test spell 1",
			"test_classes_1",
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
