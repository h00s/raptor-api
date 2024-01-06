package config

import "github.com/h00s/raptor"

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/api/v1", "HomeController", "Root"),
		raptor.Route("GET", "/api/v1/movies", "MoviesController", "Index"),
		raptor.Route("GET", "/api/v1/movies/:id", "MoviesController", "Get"),
	}
}
