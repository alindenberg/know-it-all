package userservice

import (
	"io"
	"errors"
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	var userRequest UserModels.UserRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&userRequest)
	if err != nil {
		return "", err
	}

	// Assign a random id
	user, err := userFromRequest(&userRequest)
	if err != nil {
		return "", err
	}

	// Validate group properties
	err = validateUser(user)

	if err != nil {
		return "", err
	}

	return user.UserID, UserRepository.CreateUser(user)
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

func SignIn(userId string, jsonBody io.ReadCloser) error {
	var signInRequest UserModels.UserSignInRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&signInRequest)
	if err != nil {
		return err
	}

	user, err := UserRepository.GetUserByUsername(signInRequest.Username)
	if err != nil {
		return errors.New("No user found for given username")
	}

	// Password validation
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(signInRequest.Password))
	if err != nil {
		return errors.New("Incorrect password")
	}

	return nil
}

func userFromRequest(userRequest *UserModels.UserRequest) (*UserModels.User, error) {
	userId := uuid.New().String()
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return nil, err
	}
	return &UserModels.User{userId, encryptedPassword, userRequest.Username, userRequest.Email}, nil
}
func validateUser(user *UserModels.User) error {
	return nil
}
