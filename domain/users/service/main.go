package userservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	UserRepository "github.com/alindenberg/know-it-all/domain/users/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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
	var userRequest UserModels.UserCredentials
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&userRequest)
	if err != nil {
		return "", err
	}

	// TODO: FIX FUNCTIONALITY NOW THAT USING AUTH0
	// Assign a random id
	// user, err := userFromRequest(&userRequest)
	// if err != nil {
	// 	return "", err
	// }

	// // Validate group properties
	// err = validateUser(user)

	// if err != nil {
	// 	return "", err
	// }

	// return user.UserID, UserRepository.CreateUser(user)

	return "", nil
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

func CreateUserBet(id string, jsonBody io.ReadCloser) error {
	var betRequest UserModels.UserBetRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&betRequest)
	if err != nil {
		log.Println("Error decoding bet request", err)
		return err
	}

	bet := betFromRequest(&betRequest)

	return UserRepository.CreateUserBet(id, bet)
}

func Authenticate(accessToken string) ([]string, error) {
	claims := UserModels.UserClaim{}
	tkn, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return key, nil
	})
	if err != nil || !tkn.Valid {
		return nil, err
	}

	return strings.Split(claims.Scope, " "), nil
}

func betFromRequest(request *UserModels.UserBetRequest) *UserModels.UserBet {
	return &UserModels.UserBet{
		request.MatchID,
		request.Prediction,
		false,
		false,
	}
}
func validateUser(user *UserModels.User) error {
	return nil
}
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(fmt.Sprintf("https://%s/.well-known/jwks.json", os.Getenv("appDomain")))

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = UserModels.Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
