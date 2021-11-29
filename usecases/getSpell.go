package usecases

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/JoseLuisTorrentera/academy-go-q42021/services"
	"github.com/JoseLuisTorrentera/academy-go-q42021/utils"
)

func GetSpell(name string) (*models.Spell, error) {
	spell, err := services.GetSpellByName(name)
	if err != nil {
		return nil, err
	}

	err = utils.UpdateSpellsCSV("./commons/dnd-spells.csv", spell)
	if err != nil {
		return nil, err
	}

	return spell, nil
}
