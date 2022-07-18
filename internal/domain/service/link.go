package service

import "github.com/bekzod003/link-clean/internal/domain/entities"

type linkStorage interface {
	Create(link *entities.Link) error
	Get(id int) (*entities.Link, error)
	GetByURL(url string) (*entities.Link, error)
	GetByUser(userID int) ([]*entities.Link, error)
	GetByTag(tagID int) ([]*entities.Link, error)
	Update(link *entities.Link) error
	Delete(id int) error
}

type LinkService struct {
	linkStorage linkStorage
}

func (s *LinkService) GetAllLinks(userID int64) error {
	return nil
}
