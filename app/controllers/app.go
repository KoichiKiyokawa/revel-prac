package controllers

import (
	"github.com/revel/revel"
)

// App struct
type App struct {
	*revel.Controller
}

// Index is a controller for list
func (c App) Index() revel.Result {
	return c.Render()
}
