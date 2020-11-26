package store

import "github.com/dominik-domanski/go-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindAll(start, count int) ([]model.User, error)
	FindByEmail(string) (*model.User, error)
}
