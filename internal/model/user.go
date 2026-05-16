// Package model
package model

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type User struct {
	ID           int
	FullName     string
	Email        string
	PasswordHash string
	Phone        string
	Role         UserRole
	CreatedAt    time.Time
}
