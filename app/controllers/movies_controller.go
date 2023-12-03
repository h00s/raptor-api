package controllers

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-api/app/models"
)

type MoviesController struct {
	raptor.Controller
}

func (mc *MoviesController) Index(c *raptor.Context) error {
	movies := []models.Movie{
		{
			Id:          1,
			Title:       "The Shawshank Redemption",
			Description: "A gripping tale of hope and redemption as a man finds solace and purpose in the confines of a prison.",
			Year:        1994,
		},
		{
			Id:          2,
			Title:       "The Godfather",
			Description: "A powerful crime family saga that follows the transformation of Michael Corleone from a reluctant family outsider to a ruthless mafia boss.",
			Year:        1972,
		},
	}

	return c.JSON(movies)
}
