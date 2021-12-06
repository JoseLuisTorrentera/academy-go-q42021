package usecases

import (
	"errors"
	"testing"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var spell = models.Spell{
	ID:      1,
	Name:    "test_spell",
	Classes: "test_classes",
	Level:   0,
	School:  "test_school",
}

type mockSpellApiService struct {
	mock.Mock
}

func (m mockSpellApiService) GetSpellByName(name string) (*models.Spell, error) {
	args := m.Called()
	return args.Get(0).(*models.Spell), args.Error(1)
}

type mockUpdateCSV struct {
	mock.Mock
}

func (m mockUpdateCSV) UpdateSpellsCSV(file string, spell *models.Spell) error {
	args := m.Called()
	return args.Error(0)
}

func TestUcGetSpell_GetSpell(t *testing.T) {
	testCases := []struct {
		name      string
		response  *models.Spell
		hasError  bool
		errorIs   error
		spellName string
	}{
		{
			name:      "succesful request",
			response:  &spell,
			hasError:  false,
			errorIs:   nil,
			spellName: "test_spell",
		},
		{
			name:      "failed request, not founded",
			response:  nil,
			hasError:  true,
			errorIs:   errors.New("Spell name not founded"),
			spellName: "non_spell",
		},
		{
			name:      "failed csv update",
			response:  nil,
			hasError:  true,
			errorIs:   errors.New("Problem uploading spell to csv"),
			spellName: "test_spell",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockApiService := mockSpellApiService{}
			mockApiService.On("GetSpellByName").Return(tc.response, tc.errorIs)

			mockCsvUpdate := mockUpdateCSV{}
			mockCsvUpdate.On("UpdateSpellsCSV").Return(tc.errorIs)

			usecase := NewUCGetSpell(mockApiService, mockCsvUpdate)
			spell, err := usecase.GetSpell(tc.spellName)

			assert.EqualValues(t, tc.response, spell)
			if tc.hasError {
				assert.EqualError(t, err, tc.errorIs.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
