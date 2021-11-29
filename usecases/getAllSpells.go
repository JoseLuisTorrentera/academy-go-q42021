package usecases

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/JoseLuisTorrentera/academy-go-q42021/repository"
)

type getSpells interface {
	GetAllSpells() ([]*models.Spell, error)
}

type UcGetSpells struct {
	repo getSpells
}

func NewUCGetSpells(repo repository.SpellRepo) UcGetSpells {
	return UcGetSpells{repo: repo}
}

func (uc UcGetSpells) GetAllSpells() ([]*models.Spell, error) {
	spells, err := uc.repo.GetAllSpells()
	if err != nil {
		return nil, err
	}

	return spells, nil
}
