package models

import "github.com/h00s/raptor"

type Movie struct {
	raptor.Model
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}
