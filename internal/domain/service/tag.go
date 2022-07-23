package service

import (
	"context"
	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type TagStorage interface {
	Create(ctx context.Context, tag *entities.CreateTag) (int64, error)
	Get(ctx context.Context, id int64) (*entities.Tag, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error)
	Update(ctx context.Context, tag *entities.UpdateTag) error
	Delete(ctx context.Context, id int64) error
}

type tagService struct {
	storage TagStorage
}

func NewTagService(storage TagStorage) *tagService {
	return &tagService{
		storage: storage,
	}
}

func (t *tagService) Create(tag *entities.CreateTag) (int64, error) {
	return t.storage.Create(context.TODO(), tag)
}

func (t *tagService) Get(id int64) (*entities.Tag, error) {
	return t.storage.Get(context.TODO(), id)
}

func (t *tagService) GetByUser(userID int64) ([]*entities.Tag, error) {
	return t.storage.GetByUser(context.TODO(), userID)
}

func (t *tagService) Update(tag *entities.UpdateTag) error {
	return t.storage.Update(context.TODO(), tag)
}
