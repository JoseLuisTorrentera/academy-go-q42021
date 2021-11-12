package repository

import (
	"encoding/csv"
	"os"
	
)

func GetAllData() ([]) {
	csvfile, err := os.Open("dnd-spells.csv")
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvfile)

}
