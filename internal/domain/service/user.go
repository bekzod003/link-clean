package service

import "github.com/bekzod003/link-clean/internal/domain/entities"

type UserStorage interface {
	Create(user *entities.User) error
	Get(id int) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
}
