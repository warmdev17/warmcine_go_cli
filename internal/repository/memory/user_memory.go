package memory

import (
	"errors"

	"github.com/warmdev17/warmcine_go_cli/internal/model"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type UserMemoryRepository struct {
	users  []model.User
	nextID int
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users:  make([]model.User, 0),
		nextID: 1,
	}
}

func (r *UserMemoryRepository) Create(user model.User) (model.User, error) {
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return model.User{}, ErrEmailAlreadyExists
		}
	}

	user.ID = r.nextID
	r.nextID++

	r.users = append(r.users, user)

	return user, nil
}

func (r *UserMemoryRepository) FindByEmail(email string) (model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return model.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) FindByID(id int) (model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return model.User{}, ErrUserNotFound
}

func (r *UserMemoryRepository) FindAll() ([]model.User, error) {
	return r.users, nil
}
