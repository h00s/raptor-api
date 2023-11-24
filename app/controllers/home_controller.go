package controllers

import (
	"github.com/h00s/raptor"
)

type HomeController struct {
	raptor.Controller
}

func NewHomeController() *HomeController {
	hc := &HomeController{}
	hc.Name = "Home"
	hc.RegisterActions(
		raptor.Action("Index", hc.Index),
		raptor.Action("Example", hc.Example),
	)
	return hc
}

func (hc *HomeController) Index(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Index")
	json := raptor.Map{
		"controller": "Home",
		"action":     "Index",
	}
	return c.JSON(json)
}

func (hc *HomeController) Example(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Example")
	json := raptor.Map{
		"controller": "Home",
		"action":     "Example",
	}
	return c.JSON(json)
}
