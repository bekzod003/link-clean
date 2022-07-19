package service

import (
	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type userStorage interface {
	Create(user *entities.User) (int64, error)
	Update(user *entities.User) error
	Get(id int64) (*entities.User, error)
}

type userService struct {
	storage userStorage
	log     logger.LoggerI
}

func NewUserService(storage userStorage) *userService {
	return &userService{
		storage: storage,
	}
}

func (s *userService) Create(user *entities.User) (int64, error) {
	s.log.Info("Create user", logger.Any("user", user))
	return s.storage.Create(user)
}

func (s *userService) Update(user *entities.User) error {
	s.log.Info("Update user", logger.Any("user", user))
	return s.storage.Update(user)
}

func (s *userService) Get(id int64) (*entities.User, error) {
	s.log.Info("Get user", logger.Any("id", id))
	return s.storage.Get(id)
}
