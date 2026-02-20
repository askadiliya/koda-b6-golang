package modules

import (
	"errors"
	"koda-b6-weekly1-go/models"
)

type CartModule struct {
	Items []models.CartItem
}

func NewCartModule() *CartModule {
	return &CartModule{
		Items: []models.CartItem{},
	}
}

func (c *CartModule) AddToCart(menu models.Menu) error {
	if !menu.Available {
		return errors.New("menu tidak tersedia")
	}

	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Menu.ID == menu.ID {
			c.Items[i].Qty++
			c.Items[i].SubTotal = c.Items[i].Menu.Price * c.Items[i].Qty
			return nil
		}
	}

	newItem := models.CartItem{
		Menu:     menu,
		Qty:      1,
		SubTotal: menu.Price,
	}

	c.Items = append(c.Items, newItem)
	return nil
}

func (c *CartModule) GetTotal() int {
	total := 0

	for i := 0; i < len(c.Items); i++ {
		total += c.Items[i].SubTotal
	}

	return total
}

func (c *CartModule) ClearCart() {
	c.Items = []models.CartItem{}
}

func (c *CartModule) IsEmpty() bool {
	return len(c.Items) == 0
}