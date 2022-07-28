package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/mazen160/go-random"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

func createTag(ctx context.Context, storage *tagStorage) (tagID, userID int64, err error) {
	title, err := random.String(6)
	if err != nil {
		return 0, 0, err
	}

	userID, err = createUser(ctx, NewUserStorage(newClient()))
	if err != nil {
		return 0, 0, err
	}

	tagID, err = storage.Create(ctx, &entities.CreateTag{
		Title:  title,
		UserID: userID,
	})

	return tagID, userID, err
}

func createTags(ctx context.Context, storage *tagStorage) (userID int64, err error) {
	userID, err = createUser(ctx, NewUserStorage(newClient()))
	if err != nil {
		return 0, err
	}

	for i := 0; i < 10; i++ {
		title, err := random.String(6)
		if err != nil {
			return 0, err
		}

		_, err = storage.Create(ctx, &entities.CreateTag{
			Title:  title,
			UserID: userID,
		})
	}

	return userID, err
}

func TestTagStorage_Create(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewTagStorage(newClient())

	tagID, _, err := createTag(ctx, storage)

	if err != nil {
		t.Error("Error while creating tag", err)
	}

	if tagID == 0 {
		t.Error("Error while creating user, userID is 0")
	}
}

func TestTagStorage_GetByUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewTagStorage(newClient())

	userID, err := createTags(ctx, storage)
	if err != nil {
		t.Error("Error while creating tag", err)
	}

	tags, err := storage.GetByUser(ctx, userID)
	if err != nil {
		t.Error("Error while getting tag by user", err)
	}

	for _, tag := range tags {
		if tag.UserID != userID {
			t.Error("Tag user id does not match users id")
		}
	}
}

func TestTagStorage_Get(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewTagStorage(newClient())

	tagID, userID, err := createTag(ctx, storage)
	if err != nil {
		t.Error("Error while creating tag", err)
	}

	tag, err := storage.Get(ctx, tagID)
	if err != nil {
		t.Error("Error while getting tag", err)
	}

	if tag == nil {
		t.Error("Returned tag is nil")
	}

	if tag.UserID != userID {
		t.Error("User id of tag and created user id does not match")
	}
}

func TestTagStorage_Update(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewTagStorage(newClient())

	tagID, _, err := createTag(ctx, storage)
	if err != nil {
		t.Error("Error while creating tag", err)
	}

	if err = storage.Update(
		ctx,
		&entities.UpdateTag{
			ID:    tagID,
			Title: time.Now().String(),
		},
	); err != nil {
		t.Error("Error while updating tag", err)
	}
}

func TestTagStorage_Delete(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	storage := NewTagStorage(newClient())

	tagID, _, err := createTag(ctx, storage)
	if err != nil {
		t.Error("Error while creating tag", err)
	}

	if err = storage.Delete(ctx, tagID); err != nil {
		t.Error("Error while deleting tag")
	}

	tag, err := storage.Get(ctx, tagID)

	if tag != nil {
		t.Error("Deleted tag is returned")
	}

	if err == nil {
		t.Error("No error, but should be no rows in result set")
	}
}
