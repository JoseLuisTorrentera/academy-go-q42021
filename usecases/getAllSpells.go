package usecases

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type getSpells interface {
	GetSpellsFromCsv() ([]*models.Spell, error)
}

type UcGetSpells struct {
	repo getSpells
}

func NewUCGetSpells(repo getSpells) UcGetSpells {
	return UcGetSpells{repo: repo}
}

func (uc UcGetSpells) GetAllSpells() ([]*models.Spell, error) {
	spells, err := uc.repo.GetSpellsFromCsv()
	if err != nil {
		return nil, err
	}

	return spells, nil
}
