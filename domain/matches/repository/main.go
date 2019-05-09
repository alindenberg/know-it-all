package matchesrepository

import (
	"fmt"
	"errors"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	mongo "github.com/alindenberg/know-it-all/database"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
)

var COLLECTION = "matches"

func GetAllMatches() ([]*MatchModels.Match, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*MatchModels.Match
	for cur.Next(context.TODO()) {
		var elem MatchModels.Match
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetMatch(id string) (*MatchModels.Match, error) {
	var collection = mongo.GetDbClient().Collection(COLLECTION)
	result := MatchModels.Match{}
	err := collection.FindOne(context.TODO(), bson.D{{"matchid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateMatch(match MatchModels.Match) error {
	_, err := mongo.GetDbClient().Collection(COLLECTION).InsertOne(context.TODO(), match)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMatch(id string) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"matchid", id}})

	if err != nil {
		return err
	}

	if(result.DeletedCount == 0) {
		return errors.New(fmt.Sprintf("Document with id %s was not found", id))
	}
	return nil
}

func ResolveMatch(id string, matchResult *MatchModels.MatchResult) error {
	update := bson.D{
		{"$set", bson.D{
			{"homescore", matchResult.HomeScore},
			{"awayscore", matchResult.AwayScore},
		}},
	}
	_, err := mongo.GetDbClient().Collection(COLLECTION).UpdateOne(context.TODO(), bson.D{{"matchid", id}}, update)

	return err
}
