package controllers

import (
	"fmt"
	"net/http"

	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/models"
	"github.com/h00s/raptor-api/app/services"
)

type MoviesController struct {
	raptor.Controller

	Ms *services.MoviesService
}

func (mc *MoviesController) Index(c *raptor.Context) error {
	return c.JSON(mc.Ms.All())
}

func (mc *MoviesController) Get(c *raptor.Context) error {
	id, err := c.ParamsInt("id")
	if err == nil {
		if movie := mc.Ms.Find(id); movie != nil {
			return c.JSON(movie)
		}
	}

	return c.JSON(models.NewError(http.StatusNotFound, fmt.Sprintf("Movie with id %d not found", id)), http.StatusNotFound)
}
