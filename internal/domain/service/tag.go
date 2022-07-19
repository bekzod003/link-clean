package service

import (
	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type TagStorage interface {
	Create(tag *entities.CreateTag) (int64, error)
	Get(id int64) (*entities.Tag, error)
	GetByUser(userID int64) ([]*entities.Tag, error)
	Update(tag *entities.UpdateTag) error
	Delete(id int64) error
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
	return t.storage.Create(tag)
}

func (t *tagService) Get(id int64) (*entities.Tag, error) {
	return t.storage.Get(id)
}

func (t *tagService) GetByUser(userID int64) ([]*entities.Tag, error) {
	return t.storage.GetByUser(userID)
}

func (t *tagService) Update(tag *entities.UpdateTag) error {
	return t.storage.Update(tag)
}
