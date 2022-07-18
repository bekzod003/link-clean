package service

import "github.com/bekzod003/link-clean/internal/domain/entities"

type UserStorage interface {
	Create(user *entities.User) (int64, error)
	Update(user *entities.User) error
	Get(id int64) (*entities.User, error)
}
