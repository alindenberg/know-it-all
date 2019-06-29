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
		[]UserModels.UserBet{},
		[]string{},
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

func AddFriend(userId string, friendId string) error {
	return UserRepository.AddFriend(userId, friendId)
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
		for _, bet := range user.Bets {
			if bet.MatchID == matchID {
				// wonBet := false
				bet.IsResolved = true
				if bet.Prediction == correctPrediction {
					bet.Won = true
					// wonBet = true
				}
				go UserRepository.UpdateUserBet(user.UserID, &bet)
				// go LeaderboardService.UpdateLeaderboard(userId, wonBet)
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

func getCorrectPrediction(homeScore int, awayScore int) UserModels.Prediction {
	if homeScore == awayScore {
		return UserModels.Draw
	} else if homeScore > awayScore {
		return UserModels.HomeTeam
	}
	return UserModels.AwayTeam
}
