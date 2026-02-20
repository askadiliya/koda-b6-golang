package models

type CartItem struct {
	Menu     Menu
	Qty      int
	SubTotal int
}