package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return []raptor.Route{
		raptor.NewRoute("GET", "/api/v1/home", "Home", "Index"),
		raptor.NewRoute("GET", "/api/v1/home/example", "Home", "Example"),
	}
}
