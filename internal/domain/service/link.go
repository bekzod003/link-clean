package service

import (
	"context"
	"github.com/bekzod003/link-clean/internal/domain/entities"
)

type linkStorage interface {
	Create(ctx context.Context, link *entities.Link) error
	Get(ctx context.Context, id int) (*entities.Link, error)
	GetByUser(ctx context.Context, userID int) ([]*entities.Link, error)
	GetByTag(ctx context.Context, tagID int) ([]*entities.Link, error)
	Update(ctx context.Context, link *entities.Link) error
	Delete(ctx context.Context, id int) error
}

type linkService struct {
	linkStorage linkStorage
}

func NewLinkService(linkStorage linkStorage) *linkService {
	return &linkService{
		linkStorage: linkStorage,
	}
}
