package leaguerepository

import (
	"context"
	"errors"
	"fmt"

	mongo "github.com/alindenberg/know-it-all/database"
	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION = "leagues"

func GetAllLeagues() ([]*LeagueModels.League, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*LeagueModels.League
	for cur.Next(context.TODO()) {
		var elem LeagueModels.League
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetLeague(id string) (*LeagueModels.League, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result := LeagueModels.League{}
	err := collection.FindOne(context.TODO(), bson.D{{"leagueid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateLeague(league LeagueModels.League) error {
	_, err := mongo.GetDbClient().Collection(COLLECTION).InsertOne(context.TODO(), league)
	return err
}

func CreateLeagueMatch(leagueId string, match *LeagueModels.LeagueMatch) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.D{
			{"leagueid", leagueId},
		},
		bson.D{
			{"$push", bson.D{{"matches", match}}},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteLeague(id string) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"leagueid", id}})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("Document with id %s was not found", id))
	}
	return nil
}
