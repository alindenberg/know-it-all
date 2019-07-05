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

	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	UserRepository "github.com/alindenberg/know-it-all/domain/users/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GetUser(id string) (*UserModels.User, error) {
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

	user := UserModels.User{
		userRequest.UserID,
		userRequest.Email,
		"",
		[]UserModels.UserBet{},
		[]string{},
		0,
		0,
		0.0,
	}

	return user.UserID, UserRepository.CreateUser(&user)
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

func CreateUsername(id string, jsonBody io.ReadCloser) error {
	var usernameRequest UserModels.CreateUsernameRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&usernameRequest)
	if err != nil {
		return err
	}

	err = validateUsername(usernameRequest.Username)
	if err != nil {
		return err
	}

	return UserRepository.CreateUsername(id, usernameRequest.Username)

}
func CreateUserBet(id string, jsonBody io.ReadCloser) error {
	var betRequest UserModels.CreateBetRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&betRequest)
	if err != nil {
		log.Println("Error decoding CreateBetRequest - ", err)
		return err
	}

	bet := betFromRequest(&betRequest)

	return UserRepository.CreateUserBet(id, bet)
}

func AddFriend(userId string, friendId string) error {
	return UserRepository.AddFriend(userId, friendId)
}

func UpdateUserBet(userID string, matchID string, jsonBody io.ReadCloser) error {
	_, err := uuid.Parse(matchID)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing match id : %s", err.Error()))
	}

	var updateBetRequest UserModels.UpdateBetRequest
	decoder := json.NewDecoder(jsonBody)
	err = decoder.Decode(&updateBetRequest)
	if err != nil {
		log.Println("Error decoding UpdateBetRequest -", err.Error)
		return err
	}

	return UserRepository.UpdateUserBet(userID, matchID, updateBetRequest.Prediction)
}

// Heavy load function for right now. Never done by a user or within the app,
// only by admin / bot on a daily basis. Look to refactor data model in future
func ResolveBets(matchID string, matchResult *LeagueModels.LeagueMatchResult) error {
	usersWithBets, err := UserRepository.GetUsersWithBetOnMatch(matchID)
	if err != nil {
		return err
	}

	correctPrediction := getCorrectPrediction(matchResult.HomeScore, matchResult.AwayScore)
	for _, user := range usersWithBets {
		for i := 0; i < len(user.Bets); i++ {
			if user.Bets[i].MatchID == matchID {
				if user.Bets[i].IsResolved {
					return errors.New(fmt.Sprintf("Error : User (%s) bet already resolved for match %s", user.UserID, matchID))
				} else if user.Bets[i].Prediction == correctPrediction {
					user.Bets[i].Won = true
					user.Wins = user.Wins + 1
				} else {
					user.Bets[i].Won = false
					user.Losses = user.Losses + 1
				}
				user.WinPercentage = float64(user.Wins) / float64(user.Wins+user.Losses)
				user.Bets[i].IsResolved = true
				go UserRepository.UpdateUser(user)
				break
			}
		}
	}

	return nil

	// return UserRepository.ResolveBets(matcmatchIDhId, correctPrediction)
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

func betFromRequest(request *UserModels.CreateBetRequest) *UserModels.UserBet {
	return &UserModels.UserBet{
		request.MatchID,
		request.LeagueID,
		request.Prediction,
		false,
		false,
	}
}
func validateUser(user *UserModels.User) error {
	return validateUsername(user.Username)
}
func validateUsername(username string) error {
	if len(username) < 5 || len(username) > 20 {
		return errors.New("Username must be between 5 and 20 characters in length.")
	}
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

func getCorrectPrediction(homeScore int, awayScore int) UserModels.Prediction {
	if homeScore == awayScore {
		return UserModels.Draw
	} else if homeScore > awayScore {
		return UserModels.HomeTeam
	}
	return UserModels.AwayTeam
}
