package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/controllers"
)

func Controllers() raptor.Controllers {
	hc := &controllers.HomeController{}

	return raptor.RegisterControllers(
		raptor.RegisterController("Home", &hc.Controller,
			raptor.Action("Root", hc.Root),
		),
	)
}
