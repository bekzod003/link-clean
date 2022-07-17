package v1

type handler struct {
	linkUsecase LinkUsecase
}

type LinkUsecase interface {
	GetAllLinks(userID int64) error
}

func GetAllLinks() error {
	return nil
}
