package usecases

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type getSpell interface {
	GetSpellByName(name string) (*models.Spell, error)
}

type updateCSV interface {
	UpdateSpellsCSV(file string, spell *models.Spell) error
}

type UcGetSpell struct {
	service   getSpell
	updateCSV updateCSV
}

func NewUCGetSpell(s getSpell, up updateCSV) UcGetSpell {
	return UcGetSpell{service: s, updateCSV: up}
}

func (uc UcGetSpell) GetSpell(name string) (*models.Spell, error) {
	spell, err := uc.service.GetSpellByName(name)
	if err != nil {
		return nil, err
	}

	err = uc.updateCSV.UpdateSpellsCSV("./commons/dnd-spells.csv", spell)
	if err != nil {
		return nil, err
	}

	return spell, nil
}
