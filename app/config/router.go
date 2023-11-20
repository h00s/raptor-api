package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/controllers"
)

func Router() *raptor.Router {
	hc := controllers.HomeController{}

	r := raptor.NewRouter()

	r.AddControllerRoute("Home", &hc,
		raptor.NewRoute("GET", "/api/v1/home", hc.Index),
		raptor.NewRoute("GET", "/api/v1/home/example", hc.Example),
	)

	return r
}
