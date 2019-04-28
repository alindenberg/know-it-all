package leaguerepository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	mongo "github.com/alindenberg/know-it-all/database"
	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
)

var COLLECTION = "leagues"

func GetAllLeagues() ([]*LeagueModels.League, error) {
	collection := mongo.Db.Collection(COLLECTION)
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
	collection := mongo.Db.Collection(COLLECTION)
	result := LeagueModels.League{}
	err := collection.FindOne(context.TODO(), bson.D{{"leagueid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateLeague(league LeagueModels.League) error {
	_, err := mongo.Db.Collection(COLLECTION).InsertOne(context.TODO(), league)
	if err != nil {
		return err
	}
	return nil
}