package usercontroller

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

var COLLECTION = "users"

func GetUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

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
	results, err := UserService.GetAllUsers()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := UserService.CreateUser(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func DeleteUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := UserService.DeleteUser(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
