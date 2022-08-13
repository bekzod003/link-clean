package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type userStorage interface {
	Create(ctx context.Context, user *entities.User) (int64, error)
	CreateWithGivenId(ctx context.Context, user *entities.User) error
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

func (s *userService) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	s.log.Info("Create user request", logger.Any("user", user))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	var (
		userID int64
		err    error
	)

	if user.ID == 0 {
		userID, err = s.storage.Create(ctx, user)
	} else {
		userID, err = s.storage.Create(ctx, user)
	}

	if err != nil {
		s.log.Error("Error while creating user", zap.Error(err))
		return nil, err
	}
	user.ID = userID

	s.log.Info("Successfully user has been created", zap.Any("user", user))
	return user, nil
}

func (s *userService) Update(ctx context.Context, user *entities.User) (err error) {
	s.log.Info("Update user request", logger.Any("user", user))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	if err = s.storage.Update(ctx, user); err != nil {
		s.log.Error("Error while updating user", zap.Error(err))
		return
	}

	s.log.Info("Successfully updated user")
	return
}

func (s *userService) Get(ctx context.Context, id int64) (*entities.User, error) {
	s.log.Info("Get user request", logger.Any("id", id))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	user, err := s.storage.Get(ctx, id)
	if err != nil {
		s.log.Error("Error while getting user by id", zap.Error(err))
		return nil, err
	}

	s.log.Info("Successfully updated user", zap.Any("user", user))

	return s.storage.Get(context.TODO(), id)
}
