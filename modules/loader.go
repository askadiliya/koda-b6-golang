package modules

import (
	"encoding/json"
	"errors"
	"koda-b6-weekly1-go/models"
	"os"
)

type DataStore struct {
	Categories []models.Category `json:"categories"`
	Menus      []models.Menu     `json:"menus"`
}

func LoadData(path string) (DataStore, error) {
	var store DataStore

	file, err := os.ReadFile(path)
	if err != nil {
		return store, errors.New("gagal membaca file data/menu.json")
	}

	err = json.Unmarshal(file, &store)
	if err != nil {
		return store, errors.New("format JSON tidak valid")
	}

	return store, nil
}