package repository

import (
	"github.com/rafaelarodrigues90/go-apirest/src/modules/profile/model"
)

// UserRepository interface
type UserRepository interface {
	Save(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindByID(string) (*model.User, error)
	FindAll() (model.Users, error)
}