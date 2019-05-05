package groupcontroller

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	GroupService "github.com/alindenberg/know-it-all/domain/groups/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

var COLLECTION = "groups"

func GetGroup(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	result, err := GroupService.GetGroup(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllGroups(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	results, err := GroupService.GetAllGroups()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateGroup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := GroupService.CreateGroup(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}