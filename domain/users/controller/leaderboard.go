package usercontroller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
)

func GetLeaderboard(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res, err := UserService.GetLeaderboard()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetLeaderboardForUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID := params.ByName("userId")
	res, err := UserService.GetLeaderboardForUser(userID)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
