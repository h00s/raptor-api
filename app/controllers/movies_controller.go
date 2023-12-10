package controllers

import (
	"fmt"
	"net/http"

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

func (mc *MoviesController) Get(c *raptor.Context) error {
	movies := map[int]models.Movie{
		1: {
			Id:          1,
			Title:       "The Shawshank Redemption",
			Description: "A gripping tale of hope and redemption as a man finds solace and purpose in the confines of a prison.",
			Year:        1994,
		},
		2: {
			Id:          2,
			Title:       "The Godfather",
			Description: "A powerful crime family saga that follows the transformation of Michael Corleone from a reluctant family outsider to a ruthless mafia boss.",
			Year:        1972,
		},
	}

	id, err := c.ParamsInt("id")

	if err == nil {
		if movie, ok := movies[id]; ok {
			return c.JSON(movie)
		}
	}

	return c.JSON(models.NewError(http.StatusNotFound, fmt.Sprintf("Movie with id %d not found", id)), http.StatusNotFound)
}
