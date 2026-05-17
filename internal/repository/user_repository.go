// Package repository
package repository

import "github.com/warmdev17/warmcine_go_cli/internal/model"

type UserRepository interface {
	Create(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
	FindAll() ([]model.User, error)
}
