package usercontroller

import (
	"encoding/json"
	"net/http"
	"strings"

	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
	"github.com/julienschmidt/httprouter"
)

func GetUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("userId")

	result, err := UserService.GetUser(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	queryValues := req.URL.Query()
	username := queryValues.Get("username")
	result, err := UserService.GetAllUsers(username)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := UserService.CreateUser(req.Body)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			SharedResponses.Duplicate(w, err)
		}

		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func DeleteUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("userId")

	err := UserService.DeleteUser(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUserBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	matchId := params.ByName("matchId")

	err := UserService.DeleteUserBet(userId, matchId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DeleteUserFriend(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	friendId := params.ByName("friendId")

	err := UserService.DeleteUserFriend(userId, friendId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreateUsername(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("userId")

	err := UserService.CreateUsername(id, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}

func CreateUserBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("userId")

	err := UserService.CreateUserBet(id, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func UpdateUserBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID := params.ByName("userId")
	matchID := params.ByName("matchId")

	err := UserService.UpdateUserBet(userID, matchID, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func AddFriend(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID := params.ByName("userId")
	friendID := params.ByName("friendId")

	err := UserService.AddFriend(userID, friendID)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetUserFriends(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("userId")

	result, err := UserService.GetUserFriends(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
