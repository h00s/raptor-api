package config

import "github.com/h00s/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Route("GET", "/", "HomeController", "Root"),
			raptor.Route("GET", "/movies", "MoviesController", "Index"),
			raptor.Route("GET", "/movies/:id", "MoviesController", "Get"),
		),
	)
}
