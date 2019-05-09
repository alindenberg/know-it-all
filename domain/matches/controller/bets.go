package matchcontroller

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	BetService "github.com/alindenberg/know-it-all/domain/matches/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

// func GetBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	id := params.ByName("betId")
// 	userId := params.ByName("id")
//
// 	result, err := BetService.GetBet(id, userId)
// 	if err != nil {
// 		SharedResponses.Error(w, err)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(result)
// }

func GetAllBetsForUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userId := params.ByName("id")

	results, err := BetService.GetAllBetsForUser(userId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func GetAllBetsForMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	matchId := params.ByName("id")

	results, err := BetService.GetAllBetsForMatch(matchId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userId := params.ByName("id")

	id, err := BetService.CreateBet(req.Body, userId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func DeleteBet(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("betId")
	userId := params.ByName("id")

	err := BetService.DeleteBet(id, userId)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}
