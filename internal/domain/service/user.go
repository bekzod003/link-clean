package service

import (
	"context"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type userStorage interface {
	Create(ctx context.Context, user *entities.User) (int64, error)
	Update(ctx context.Context, user *entities.User) error
	Get(ctx context.Context, id int64) (*entities.User, error)
}

type userService struct {
	storage userStorage
	log     logger.LoggerI
}

func NewUserService(storage userStorage, log logger.LoggerI) *userService {
	return &userService{
		storage: storage,
		log:     log,
	}
}

func (s *userService) Create(ctx context.Context, user *entities.User) (int64, error) {
	s.log.Info("Create user", logger.Any("user", user))
	return s.storage.Create(context.TODO(), user)
}

func (s *userService) Update(ctx context.Context, user *entities.User) error {
	s.log.Info("Update user", logger.Any("user", user))
	return s.storage.Update(context.TODO(), user)
}

func (s *userService) Get(ctx context.Context, id int64) (*entities.User, error) {
	s.log.Info("Get user", logger.Any("id", id))
	return s.storage.Get(context.TODO(), id)
}
