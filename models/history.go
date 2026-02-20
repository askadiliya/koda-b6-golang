package models

import "time"

type Transaction struct {
	ID        int
	Items     []CartItem
	Total     int
	CreatedAt time.Time
}