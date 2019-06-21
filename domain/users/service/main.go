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
	"time"

	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	UserRepository "github.com/alindenberg/know-it-all/domain/users/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func CreateUserSession(jsonBody io.ReadCloser) (*UserModels.UserKeys, error) {
	var userCredentials *UserModels.UserCredentials
	err := json.NewDecoder(jsonBody).Decode(&userCredentials)
	if err != nil {
		return nil, err
	}

	user, err := UserRepository.GetUserByUsername(userCredentials.Username)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("No user found for given username")
	}

	// Password validation
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userCredentials.Password))
	if err != nil {
		return nil, errors.New("Incorrect password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	signedAccessToken, err := getSignedToken(userCredentials.Username, "", expirationTime)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Create renewal token that is good for 1 day
	// (add remaining minutes to access token expiration time)
	renewTokenExpirationTime := expirationTime.Add(1395 * time.Minute)
	signedRenewToken, err := getSignedToken(userCredentials.Username, "", renewTokenExpirationTime)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	UserRepository.CreateUserKeys(&UserModels.UserKeys{userCredentials.Username, signedAccessToken, signedRenewToken})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &UserModels.UserKeys{AccessToken: signedAccessToken, RenewToken: signedRenewToken}, nil
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

// func RenewAuthentication()

func getSignedToken(username string, scope string, expirationTime time.Time) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &UserModels.UserClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
		Scope: scope,
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	signedToken, err := token.SignedString([]byte(os.Getenv("jwtKey")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Println(err.Error)
		return "", err
	}

	return signedToken, nil
}
func userFromRequest(userCredentials *UserModels.UserCredentials) (*UserModels.User, error) {
	userId := uuid.New().String()
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userCredentials.Password), 10)
	if err != nil {
		return nil, err
	}
	return &UserModels.User{userId, encryptedPassword, userCredentials.Username, userCredentials.Email}, nil
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
