package groupservice

import (
	"io"
	"encoding/json"
	"github.com/google/uuid"
	GroupModels "github.com/alindenberg/know-it-all/domain/groups/models"
	GroupRepository "github.com/alindenberg/know-it-all/domain/groups/repository"
)

func GetGroup(id string) (*GroupModels.Group, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return GroupRepository.GetGroup(id)
} 

func GetAllGroups() ([]*GroupModels.Group, error) {
	return GroupRepository.GetAllGroups()
}

func CreateGroup(jsonBody io.ReadCloser) (string, error) {
	var group GroupModels.Group
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&group)
	if err != nil {
		return "", err
	}

	// TODO: Validation

	id := uuid.New().String()
	group.GroupID = id

	return id, GroupRepository.CreateGroup(group)


}