package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/controllers"
	"github.com/h00s/raptor-api/app/services"
)

func Controllers() raptor.Controllers {
	ms := services.NewMoviesService()

	return raptor.RegisterControllers(
		&controllers.HomeController{},
		&controllers.MoviesController{
			Ms: ms,
		},
	)
}
