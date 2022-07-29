package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type TagStorage interface {
	Create(ctx context.Context, tag *entities.Tag) (*entities.Tag, error)
	Get(ctx context.Context, id int64) (*entities.Tag, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error)
	Update(ctx context.Context, tag *entities.UpdateTag) error
	Delete(ctx context.Context, id int64) error
}

type tagService struct {
	storage TagStorage
	log     logger.LoggerI
}

func NewTagService(storage TagStorage, log logger.LoggerI) *tagService {
	return &tagService{
		storage: storage,
		log:     log,
	}
}

// @TODO: implement tag service
func (t *tagService) Create(ctx context.Context, tag *entities.Tag) (*entities.Tag, error) {
	t.log.Info("Create tag request", zap.Any("req", tag))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	tag, err := t.storage.Create(ctx, tag)
	if err != nil {
		t.log.Error("Error while creating tag", zap.Error(err))
		return nil, err
	}

	t.log.Info("Create tag response", zap.Any("resp", tag))
	return tag, err
}

func (t *tagService) Get(ctx context.Context, id int64) (*entities.Tag, error) {
	t.log.Info("Get tag request", zap.Int64("id", id))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	tag, err := t.storage.Get(ctx, id)
	if err != nil {
		t.log.Error("Error while getting tag by id", zap.Error(err))
		return nil, err
	}

	t.log.Info("Get tag response", zap.Any("resp", tag))
	return tag, nil
}

func (t *tagService) GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error) {
	t.log.Info("Get tag by user request", zap.Int64("userID", userID))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	tags, err := t.storage.GetByUser(ctx, userID)
	if err != nil {
		t.log.Error("Error while getting tags by user", zap.Error(err))
	}

	t.log.Info("Get tag by user response", zap.Any("resp", tags))
	return tags, nil
}

func (t *tagService) Update(ctx context.Context, tag *entities.UpdateTag) (err error) {
	t.log.Info("Update tag request", zap.Any("req", tag))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	if err = t.storage.Update(ctx, tag); err != nil {
		t.log.Error("Error while updating tag", zap.Error(err))
		return
	}

	t.log.Info("Tag successfully updated", zap.Any("tag", tag))
	return
}

func (t *tagService) Delete(ctx context.Context, id int64) (err error) {
	t.log.Info("Delete tag request", zap.Int64("ID", id))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	if err = t.storage.Delete(ctx, id); err != nil {
		t.log.Error("Error while deleting tag", zap.Error(err))
		return
	}

	t.log.Info("Tag successfully deleted")
	return nil
}
