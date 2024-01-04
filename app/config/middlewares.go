package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/middlewares"
)

func Middlewares() raptor.Middlewares {
	return raptor.Middlewares{
		&middlewares.ActionLogger{},
	}
}
