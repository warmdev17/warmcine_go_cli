// Package model
package model

import "time"

type BookingStatus string

const (
	BookingPending   BookingStatus = "pending"
	BookingConfirmed BookingStatus = "confirmed"
	BookingCanceled  BookingStatus = "canceled"
)

type Booking struct {
	ID         int
	UserID     int
	ShowtimeID int
	SeatIDs    []int
	TotalPrice int
	Status     BookingStatus
	CreatedAt  time.Time
}
