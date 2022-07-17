package service

type linkStorage interface {
	// Create(link *entity.Link) error
	// Get(id int) (*Link, error)
	// GetByURL(url string) (*Link, error)
	// GetByUser(userID int) ([]*Link, error)
	// GetByTag(tagID int) ([]*Link, error)
	// Update(link *Link) error
	// Delete(id int) error
}

type LinkService struct {
	linkStorage linkStorage
}

func (s *LinkService) GetAllLinks(userID int64) error {
	return nil
}
