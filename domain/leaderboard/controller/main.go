package usercontroller

import (
	"encoding/json"
	"net/http"

	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

func GetLeaderboard(w http.ResponseWriter, req *http.Request, _ httprouter.Prams) {
	res, err := LeaderboardService.GetLeaderboard()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetLeaderboardForUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res, err := LeaderboardService.GetLeaderboardForUser()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
