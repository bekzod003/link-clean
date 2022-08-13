package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/mazen160/go-random"

	"github.com/bekzod003/link-clean/internal/domain/entities"
)

func createUser(ctx context.Context, storage *userStorage) (int64, error) {
	user, err := fillUserFields()
	if err != nil {
		return 0, err
	}

	userID, err := storage.Create(
		ctx,
		user,
	)
	return userID, err
}

func createUserWithGivenId(ctx context.Context, storage *userStorage) (int64, error) {
	user, err := fillUserFields()
	user.ID = 160
	if err != nil {
		return 0, err
	}

	err = storage.CreateWithGivenId(
		ctx,
		user,
	)
	return user.ID, err
}

func fillUserFields() (*entities.User, error) {
	userName, err := random.String(6)
	if err != nil {
		return nil, err
	}
	firstName, err := random.String(6)
	if err != nil {
		return nil, err
	}
	lastName, err := random.String(6)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		Username:  userName,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func TestUserStorage_Create(t *testing.T) {
	client := newClient()
	storage := NewUserStorage(client)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	userID, err := createUser(ctx, storage)
	if err != nil {
		t.Error("Error while creating user:", err)
	}
	if userID == 0 {
		t.Error("UserID is equal to 0!")
	}
}

func TestUserStorage_CreateWithGivenId(t *testing.T) {
	client := newClient()
	storage := NewUserStorage(client)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	userID, err := createUserWithGivenId(ctx, storage)
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

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	id, err := createUser(ctx, storage)
	if err != nil {
		t.Error("Error while creating new user:", err)
	}

	res, err := storage.Get(ctx, id)
	if err != nil {
		t.Error("Error while getting user by id:", err)
	}
	if res == nil {
		t.Error("User is nil!")
	}

	fmt.Printf("Result: %+v", res)
}

func TestUserStorage_Update(t *testing.T) {
	storage := NewUserStorage(newClient())

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	id, err := createUser(ctx, storage)
	if err != nil {
		t.Error("Error while creating user:", err)
	}

	user, err := fillUserFields()
	user.ID = id
	if err != nil {
		t.Error("Error while filling user fields:", err)
	}
	err = storage.Update(ctx, user)
	if err != nil {
		t.Error("Error while updating user:", err)
	}
}
