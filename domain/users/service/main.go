package userservice

import (
	"io"
	"encoding/json"
	"github.com/google/uuid"
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	UserRepository "github.com/alindenberg/know-it-all/domain/users/repository"
)

func GetUser(id string) (*UserModels.User, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return UserRepository.GetUser(id)
}

func GetAllUsers() ([]*UserModels.User, error) {
	return UserRepository.GetAllUsers()
}

func CreateUser(jsonBody io.ReadCloser) (string, error) {
	var user UserModels.User
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&user)
	if err != nil {
		return "", err
	}

	// Assign a random id
	id := uuid.New().String()
	user.UserID = id

	// Validate group properties
	err = validateUser(&user)

	if err != nil {
		return "", err
	}

	return id, UserRepository.CreateUser(user)
}

func DeleteUser(id string) error {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return UserRepository.DeleteUser(id)
}

func validateUser(user *UserModels.User) error {
	return nil
}
