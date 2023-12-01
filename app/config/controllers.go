package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/controllers"
)

func Controllers() raptor.Controllers {
	hc := &controllers.HomeController{}
	mc := &controllers.MoviesController{}

	return raptor.RegisterControllers(
		raptor.RegisterController("Home", &hc.Controller,
			raptor.Action("Root", hc.Root),
		),
		raptor.RegisterController("Movies", &mc.Controller,
			raptor.Action("Index", mc.Index),
		),
	)
}
