package link

import (
	"github.com/bekzod003/link-clean/internal/domain/service"
)

type LinkUsecase struct {
	linkService service.LinkService
}

func NewLinkUsecase() *LinkUsecase {
	return &LinkUsecase{}
}

func (l *LinkUsecase) GetAllLinks(userId int64) error {
	return l.linkService.GetAllLinks(userId)
}
