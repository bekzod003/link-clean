package link

import (
	"context"

	"go.uber.org/zap"

	"github.com/bekzod003/link-clean/internal/domain/entities"
	"github.com/bekzod003/link-clean/pkg/logger"
)

type linkService interface {
	Create(ctx context.Context, link *entities.Link) (*entities.Link, error)
	Get(ctx context.Context, id int64) (*entities.Link, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Link, error)
	GetByTag(ctx context.Context, tagID int64) ([]*entities.Link, error)
	Update(ctx context.Context, link *entities.Link) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type tagService interface {
	Create(ctx context.Context, tag *entities.CreateTag) (id int64, err error)
	Get(ctx context.Context, id int64) (*entities.Tag, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error)
	Update(ctx context.Context, tag *entities.UpdateTag) error
	Delete(ctx context.Context, id int64) error
}

type userService interface {
	Create(ctx context.Context, user *entities.User) (int64, error)
	Update(ctx context.Context, user *entities.User) error
	Get(ctx context.Context, id int64) (*entities.User, error)
}

type LinkUsecase struct {
	linkService linkService
	tagService  tagService
	userService userService
	log         logger.LoggerI
}

// Used only to create new link usecase
type ConstructorRequest struct {
	LinkService linkService
	TagService  tagService
	UserService userService
	Log         logger.LoggerI
}

func NewLinkUsecase(request ConstructorRequest) *LinkUsecase {
	return &LinkUsecase{
		linkService: request.LinkService,
		tagService:  request.TagService,
		userService: request.UserService,
		log:         request.Log,
	}
}

func (l *LinkUsecase) GetByUser(userId int64) ([]*entities.Link, error) {
	links, err := l.linkService.GetByUser(context.TODO(), userId)
	if err != nil {
		l.log.Error("GetByUser", logger.Any("userId", userId), logger.Error(err))
		return nil, err
	}
	return links, nil
}

func (l *LinkUsecase) GetUser(id int64) (*entities.User, error) {
	l.log.Info("Get user request:", zap.Any("req", id))
	return l.userService.Get(context.TODO(), id)
}
