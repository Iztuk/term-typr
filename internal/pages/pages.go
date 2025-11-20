package pages

import (
	"term-typr/internal/pages/menu"
	"term-typr/internal/pages/practice"
)

type Page struct {
	Menu     menu.MenuModel
	Practice practice.PracticeModel
}
