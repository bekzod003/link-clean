package service

import "github.com/bekzod003/link-clean/internal/domain/entities"

type TagStorage interface {
	Create(tag *entities.Tag) error
	Get(id int) (*entities.Tag, error)
	GetByUser(userID int) ([]*entities.Tag, error)
	Update(tag *entities.UpdateTag) error
	Delete(id int) error
}

type tagService struct {
	storage TagStorage
}

func NewTagStorage(storage TagStorage) *tagService {
	return &tagService{storage: storage}
}

func (t *tagService) Create(tag *entities.Tag) error {
	return t.storage.Create(tag)
}

func (t *tagService) Get(id int) (*entities.Tag, error) {
	return t.storage.Get(id)
}

func (t *tagService) GetByUser(userID int) ([]*entities.Tag, error) {
	return t.storage.GetByUser(userID)
}

func (t *tagService) Update(tag *entities.UpdateTag) error {
	return t.storage.Update(tag)
}
