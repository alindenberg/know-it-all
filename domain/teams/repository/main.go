package teamrepository

import (
	"context"
	"errors"
	"fmt"

	mongo "github.com/alindenberg/know-it-all/database"
	TeamModels "github.com/alindenberg/know-it-all/domain/teams/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION = "teams"

func GetAllTeams() ([]*TeamModels.Team, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*TeamModels.Team
	for cur.Next(context.TODO()) {
		var elem TeamModels.Team
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetAllTeamsForLeague(leagueID string) ([]*TeamModels.Team, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{{"leagueId", leagueID}}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*TeamModels.Team
	for cur.Next(context.TODO()) {
		var elem TeamModels.Team
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetTeam(teamID string) (*TeamModels.Team, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	team := TeamModels.Team{}
	err := collection.FindOne(context.TODO(), bson.D{{"teamID", teamID}}).Decode(&team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func CreateTeam(team *TeamModels.Team) error {
	_, err := mongo.GetDbClient().Collection(COLLECTION).InsertOne(context.TODO(), &team)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTeam(teamID string) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"teamID", teamID}})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("Document with id %s was not found", teamID))
	}
	return nil
}
