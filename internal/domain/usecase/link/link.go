package link

import (
	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type linkService interface {
	Create(link *entities.Link) error
	Get(id int64) (*entities.Link, error)
	GetByUser(userID int64) ([]*entities.Link, error)
	GetByTag(tagID int64) ([]*entities.Link, error)
	Update(link *entities.Link) error
	Delete(id int64) error
}

type LinkUsecase struct {
	linkService linkService
	log         logger.LoggerI
}

func NewLinkUsecase(ls linkService, log logger.LoggerI) *LinkUsecase {
	return &LinkUsecase{
		linkService: ls,
		log:         log,
	}
}

func (l *LinkUsecase) GetByUser(userId int64) ([]*entities.Link, error) {
	links, err := l.linkService.GetByUser(userId)
	if err != nil {
		l.log.Error("GetByUser", logger.Any("userId", userId), logger.Error(err))
		return nil, err
	}
	return links, nil
}
