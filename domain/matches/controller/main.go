package matchcontroller

import (
	"encoding/json"
	"net/http"

	MatchService "github.com/alindenberg/know-it-all/domain/matches/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	"github.com/julienschmidt/httprouter"
)

func CreateMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := MatchService.CreateMatch(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func GetMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	matchID := params.ByName("matchId")

	result, err := MatchService.GetMatch(matchID)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMatches(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	queryValues := req.URL.Query()
	leagueID := queryValues.Get("leagueId")
	excludePast := queryValues.Get("excludePast")
	results, err := MatchService.GetAllMatches(leagueID, excludePast)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func ResolveMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	matchID := params.ByName("matchId")

	err := MatchService.ResolveMatch(matchID, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}
