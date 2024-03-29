package link

import (
	"context"
	"time"

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
	Create(ctx context.Context, tag *entities.Tag) (*entities.Tag, error)
	Get(ctx context.Context, id int64) (*entities.Tag, error)
	GetByUser(ctx context.Context, userID int64) ([]*entities.Tag, error)
	Update(ctx context.Context, tag *entities.UpdateTag) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type userService interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (err error)
	Get(ctx context.Context, id int64) (*entities.User, error)
}

type UsecaseLink struct {
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

func NewLinkUsecase(request ConstructorRequest) *UsecaseLink {
	return &UsecaseLink{
		linkService: request.LinkService,
		tagService:  request.TagService,
		userService: request.UserService,
		log:         request.Log,
	}
}

func (l *UsecaseLink) CreateUser(user *User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	l.log.Info("usecase Create createUserResp request", zap.Any("request", user))
	createUserResp, err := l.userService.Create(
		ctx,
		&entities.User{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	)
	if err != nil {
		l.log.Error("Error while creating createUserResp", zap.Error(err))
		return nil, err
	}

	user.ID = createUserResp.ID
	user.CreatedAt = createUserResp.CreatedAt

	return user, nil
}

func (l *UsecaseLink) GetByUser(userId int64) ([]*entities.Link, error) {
	links, err := l.linkService.GetByUser(context.TODO(), userId)
	if err != nil {
		l.log.Error("GetByUser", logger.Any("userId", userId), logger.Error(err))
		return nil, err
	}
	return links, nil
}

func (l *UsecaseLink) GetUser(id int64) (*entities.User, error) {
	l.log.Info("Get user request:", zap.Any("req", id))
	return l.userService.Get(context.TODO(), id)
}
