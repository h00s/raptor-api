package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.NewControllers(
		&controllers.NewHomeController().Controller,
	)
}
