package main

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/config"
)

func main() {
	r := raptor.NewRaptor()

	r.Controllers(config.Controllers())
	r.Routes(config.Routes())

	r.Listen()
}
