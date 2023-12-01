package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/api/v1", "Home", "Root"),
		raptor.Route("GET", "/api/v1/movies", "Movies", "Index"),
	}
}
