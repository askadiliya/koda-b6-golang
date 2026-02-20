package modules

import (
	"koda-b6-weekly1-go/models"
	"strings"
)

type MenuModule struct {
	Categories []models.Category
	Menus      []models.Menu
}

func NewMenuModule(categories []models.Category, menus []models.Menu) *MenuModule {
	return &MenuModule{
		Categories: categories,
		Menus:      menus,
	}
}

func (m *MenuModule) GetCategoryName(categoryID int) string {
	for i := 0; i < len(m.Categories); i++ {
		if m.Categories[i].ID == categoryID {
			return strings.ToUpper(m.Categories[i].Name)
		}
	}
	return "UNKNOWN"
}

func (m *MenuModule) GetMenusByCategory(categoryID int) []models.Menu {
	result := []models.Menu{}

	for i := 0; i < len(m.Menus); i++ {
		if m.Menus[i].CategoryID == categoryID {
			result = append(result, m.Menus[i])
		}
	}

	return result
}

func (m *MenuModule) GetMenuByID(menuID int) (models.Menu, bool) {
	for i := 0; i < len(m.Menus); i++ {
		if m.Menus[i].ID == menuID {
			return m.Menus[i], true
		}
	}
	return models.Menu{}, false
}