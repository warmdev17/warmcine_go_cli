// Package model
package model

type Movie struct {
	ID          int
	Title       string
	Description string
	Duration    int // minutes
	AgeRating   string
	Genre       string
}
