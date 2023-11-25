package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/api/v1/home", "Home", "Index"),
		raptor.Route("GET", "/api/v1/home/example", "Home", "Example"),
	}
}