// Package model
package model

type SeatType string

const (
	SeatNormal SeatType = "normal"
	SeatVIP    SeatType = "vip"
)

type Seat struct {
	ID     int
	RoomID int
	Row    string
	Number int
	Type   SeatType
}
