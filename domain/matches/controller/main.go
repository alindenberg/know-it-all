package matchcontroller

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	MatchService "github.com/alindenberg/know-it-all/domain/matches/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

func GetMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	result, err := MatchService.GetMatch(id)

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMatches(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	result, err := MatchService.GetAllMatches()

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateMatch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	id, err := MatchService.CreateMatch(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func ResolveMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := MatchService.ResolveMatch(id, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}

func DeleteMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := MatchService.DeleteMatch(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}
