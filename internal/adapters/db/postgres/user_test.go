package postgres

import (
	"context"
	"fmt"
	"github.com/bekzod003/link-clean/internal/domain/entities"
	"testing"
	"time"
)

// @TODO: create user with random values!
func createUser(ctx context.Context, storage *userStorage) (int64, error) {
	userID, err := storage.Create(
		ctx,
		&entities.User{
			ID:        7,
			Username:  "username",
			FirstName: "firstname",
			LastName:  "lastname",
		},
	)
	return userID, err

}

func TestUserStorage_Create(t *testing.T) {
	client := newClient()
	storage := NewUserStorage(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID, err := createUser(ctx, storage)
	if err != nil {
		t.Error("Error while creating user:", err)
	}
	if userID == 0 {
		t.Error("UserID is equal to 0!")
	}
}

func TestUserStorage_Get(t *testing.T) {
	client := newClient()
	storage := NewUserStorage(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := storage.Get(ctx, 5)
	if err != nil {
		t.Error("Error while getting user by id:", err)
	}
	if res == nil {
		t.Error("User is nil!")
	}

	fmt.Printf("Result: %+v", res)
}
