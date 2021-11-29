package services

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type Response struct {
	Name    string  `json:"name"`
	Level   int     `json:"level"`
	School  School  `json:"school"`
	Classes []Class `json:"classes"`
}

type School struct {
	Name string `json:"name"`
}

type Class struct {
	Name string `json:"name"`
}

func GetSpellByName(name string) (*models.Spell, error) {

	name = strings.ToLower(name)
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/spells/%s", name)
	response, err := http.Get(url)
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Spell not founded!")
	}

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	classes := ""
	for _, class := range responseObject.Classes {
		classes += (class.Name + ",")
	}

	newId, err := GenerateSpellIndex()
	if err != nil {
		return nil, err
	}

	spell := models.Spell{
		ID:      newId,
		Name:    responseObject.Name,
		Classes: classes,
		Level:   responseObject.Level,
		School:  responseObject.School.Name,
	}

	return &spell, nil
}

func GenerateSpellIndex() (int, error) {
	csvFile, err := os.Open("./commons/dnd-spells.csv")
	defer csvFile.Close()
	if err != nil {
		return 0, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return 0, err
	}

	id := csvLines[len(csvLines)-1][0]
	newId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	csvFile.Close()
	return newId + 1, nil
}
