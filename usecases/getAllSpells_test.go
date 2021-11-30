package usecases

import (
	"errors"
	"testing"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type mockSpellsRepo struct {
	mock.Mock
}

func (mr mockSpellsRepo) GetSpellsFromCsv() ([]*models.Spell, error) {
	arg := mr.Called()
	return arg.Get(0).([]*models.Spell), arg.Error(1)
}

func TestUcGetSpells_GetAllSpells(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       []*models.Spell
		hasError       bool
		errorIs        error
	}{
		{
			"get all spells",
			2,
			spells,
			false,
			nil,
		},
		{
			"error in repository",
			0,
			nil,
			true,
			errors.New("Error in repository"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockSpellsRepo{}
			mock.On("GetSpellsFromCsv").Return(tc.response, tc.errorIs)
			usecase := NewUCGetSpells(mock)

			spells, err := usecase.GetAllSpells()

			assert.EqualValues(t, tc.expectedLength, len(spells))
			if tc.hasError {
				assert.EqualError(t, err, tc.errorIs.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
