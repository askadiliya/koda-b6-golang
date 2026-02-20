package models

type Menu struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}