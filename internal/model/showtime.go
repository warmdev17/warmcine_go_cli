// Package model
package model

import "time"

type Showtime struct {
	ID        int
	MovieID   int
	RoomID    int
	StartTime time.Time
	EndTime   time.Time
	Price     int
}
