package service

import (
	"context"
	"time"

	"go.uber.org/zap"

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
	storage linkStorage
	log     logger.LoggerI
}

func NewLinkService(linkStorage linkStorage, log logger.LoggerI) *linkService {
	return &linkService{
		storage: linkStorage,
		log:     log,
	}
}

func (ls *linkService) Create(ctx context.Context, link *entities.Link) (*entities.Link, error) {
	ls.log.Info("Link service create request", zap.Any("req", link))

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	link, err := ls.storage.Create(ctx, link)
	if err != nil {
		ls.log.Error("Error while creating link", zap.Error(err))
		return nil, err
	}

	ls.log.Info("Link create response", zap.Any("resp", link))
	return link, nil
}

func (ls *linkService) Get(ctx context.Context, id int64) (*entities.Link, error) {
	ls.log.Info("Get link request", zap.Int64("id", id))

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	link, err := ls.storage.Get(ctx, id)

	if err != nil {
		ls.log.Error("Error while getting link", zap.Error(err))
		return nil, err
	}

	ls.log.Info("Get link response", zap.Any("resp", link))
	return link, nil
}

func (ls *linkService) GetByUser(ctx context.Context, userID int64) ([]*entities.Link, error) {
	ls.log.Info("Get link by user request", zap.Int64("id", userID))

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	links, err := ls.storage.GetByUser(ctx, userID)
	if err != nil {
		ls.log.Error("Error while getting link by user", zap.Error(err))
		return nil, err
	}

	ls.log.Info("Get link by user response", zap.Any("resp", links))
	return links, nil
}

func (ls *linkService) GetByTag(ctx context.Context, tagID int64) ([]*entities.Link, error) {
	ls.log.Info("Get link by tag request", zap.Int64("id", tagID))

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	links, err := ls.storage.GetByTag(ctx, tagID)
	if err != nil {
		ls.log.Error("Error while getting link by tag", zap.Error(err))
		return nil, err
	}

	ls.log.Info("Get link by tag response", zap.Any("resp", links))
	return links, nil
}

func (ls *linkService) Update(ctx context.Context, link *entities.Link) (err error) {
	ls.log.Info("Update link request", zap.Any("req", link))

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err = ls.storage.Update(ctx, link); err != nil {
		ls.log.Error("Error while updating link", zap.Error(err))
		return
	}

	ls.log.Info("Link successfully updated")
	return nil
}

func (ls *linkService) Delete(ctx context.Context, id int64) (err error) {
	ls.log.Info("Delete link request", zap.Int64("id", id))

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err = ls.storage.Delete(ctx, id); err != nil {
		ls.log.Error("Error while deleting link", zap.Error(err))
		return
	}

	ls.log.Info("Link successfully deleted")
	return nil
}
