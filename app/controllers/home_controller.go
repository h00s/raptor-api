package controllers

import (
	"github.com/h00s/raptor"
)

type HomeController struct {
	raptor.Controller
}

func (hc *HomeController) Root(c *raptor.Context) error {
	json := raptor.Map{
		"controller": "Home",
		"action":     "Index",
		"message":    "Hello from HomeController.Index",
	}
	return c.JSON(json)
}
