package usecases

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/JoseLuisTorrentera/academy-go-q42021/repository"
)

func GetAllSpells() ([]*models.Spell, error) {
	spells, err := repository.GetAllSpells()
	if err != nil {
		return nil, err
	}

	return spells, nil
}
