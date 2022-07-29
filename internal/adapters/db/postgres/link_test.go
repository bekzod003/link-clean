package postgres

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mazen160/go-random"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

func createLink(ctx context.Context, storage *linkStorage) (link *entities.Link, err error) {
	link = &entities.Link{}
	link.TagID, link.UserID, err = createTag(ctx, NewTagStorage(newClient()))
	if err != nil {
		return
	}

	title, err := random.String(6)
	if err != nil {
		return
	}

	description, err := random.String(6)
	if err != nil {
		return
	}

	URL, err := random.String(6)
	if err != nil {
		return
	}

	link, err = storage.Create(
		ctx,
		&entities.Link{
			URL:         URL,
			Title:       title,
			Description: description,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			UserID:      link.UserID,
			TagID:       link.TagID,
		},
	)

	return
}

func compareLinks(respLink, link *entities.Link) error {
	if respLink.ID != link.ID ||
		respLink.UserID != link.UserID ||
		respLink.TagID != link.TagID ||
		respLink.Title != link.Title ||
		respLink.URL != link.URL {
		return errors.New("Structures are not the same")
	}

	return nil
}

func TestLinkStorage_Create(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	if link.ID == 0 {
		t.Error("Error while creating link, link id is 0")
	}
}

func TestLinkStorage_Get(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	respLink, err := storage.Get(ctx, link.ID)
	if err != nil {
		t.Error("Error while getting link:", err)
	}

	if err = compareLinks(respLink, link); err != nil {
		t.Error("Error while comparing links:", err)
	}
}

func TestLinkStorage_GetByUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	returnedLinks, err := storage.GetByUser(ctx, link.UserID)
	if err != nil {
		t.Error("Error while getting link by user")
	}

	if returnedLinks == nil {
		t.Error("Returned link by user id is nil")
	}

	if len(returnedLinks) == 0 {
		t.Error("Returned link length is 0")
	}

	if err = compareLinks(returnedLinks[0], link); err != nil {
		t.Error("Links are not the same")
	}
}

func TestLinkStorage_GetByTag(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	links, err := storage.GetByTag(ctx, link.TagID)
	if err != nil {
		t.Error("Error while getting link by tag:", err)
	}

	if links == nil || len(links) == 0 {
		t.Error("Error while getting links, link is nil, or length is equals to 0")
	}
}

func TestLinkStorage_Update(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	if link == nil {
		t.Error("Error while creating link, returned link is nil")
	}

	link.Title = time.Now().String()
	link.Description = time.Now().String()

	if err = storage.Update(ctx, link); err != nil {
		t.Error("Error while updating link:", err)
	}
}

func TestLinkStorage_Delete(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewLinkStorage(newClient())
	link, err := createLink(ctx, storage)
	if err != nil {
		t.Error("Error while creating link:", err)
	}

	if err = storage.Delete(ctx, link.ID); err != nil {
		t.Error("Error while deleting link:", err)
	}
}
