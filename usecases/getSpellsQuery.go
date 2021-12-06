package usecases

import (
	"errors"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type getSpellsQuery interface {
	GetSpellsByQuery(itemType int, numItems int, numItemsWorker int) ([]*models.Spell, error)
}

type UcGetSpellsQuery struct {
	service getSpellsQuery
}

func NewUCGetSpellsQuery(service getSpellsQuery) UcGetSpellsQuery {
	return UcGetSpellsQuery{service: service}
}

func (uc UcGetSpellsQuery) GetSpellsByQuery(itemType string, numItems int, numItemsWorker int) ([]*models.Spell, error) {
	itType := 0
	switch itemType {
	case "odd":
		itType = 0
	case "even":
		itType = 1
	default:
		return nil, errors.New("Only accept odd or even")
	}
	spells, err := uc.service.GetSpellsByQuery(itType, numItems, numItemsWorker)
	if err != nil {
		return nil, err
	}

	return spells, nil
}
