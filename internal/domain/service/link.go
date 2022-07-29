package service

import (
	"context"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type linkStorage interface {
	Create(ctx context.Context, link *entities.Link) (*entities.Link, error)
	Get(ctx context.Context, id int64) (*entities.Link, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Link, error)
	GetByTag(ctx context.Context, tagID int64) ([]*entities.Link, error)
	Update(ctx context.Context, link *entities.Link) error
	Delete(ctx context.Context, id int64) error
}

type linkService struct {
	linkStorage linkStorage
	log         logger.LoggerI
}

func NewLinkService(linkStorage linkStorage, log logger.LoggerI) *linkService {
	return &linkService{
		linkStorage: linkStorage,
		log:         log,
	}
}

// @TODO: implement link service

func (ls *linkService) Create(ctx context.Context, link *entities.Link) error {
	return nil
}

func (ls *linkService) Get(ctx context.Context, id int64) (*entities.Link, error) {
	return nil, nil
}
func (ls *linkService) GetByUser(ctx context.Context, userID int64) ([]*entities.Link, error) {
	return nil, nil
}
func (ls *linkService) GetByTag(ctx context.Context, tagID int64) ([]*entities.Link, error) {
	return nil, nil
}
func (ls *linkService) Update(ctx context.Context, link *entities.Link) error {
	return nil
}
func (ls *linkService) Delete(ctx context.Context, id int64) error {
	return nil
}
