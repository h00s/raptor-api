package controllers

import (
	"github.com/h00s/raptor"
)

type HomeController struct {
	raptor.DefaultController
}

func (hc *HomeController) Index(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Index")
	json := map[string]interface{}{
		"controller": "Home",
		"action":     "Index",
	}
	return c.JSON(json)
}

func (hc *HomeController) Example(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Example")
	json := map[string]interface{}{
		"controller": "Home",
		"action":     "Example",
	}
	return c.JSON(json)
}
