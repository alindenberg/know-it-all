package matchesrepository

import (
	"context"
	"log"

	mongo "github.com/alindenberg/know-it-all/database"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION = "matches"

func GetAllMatches() []*MatchModels.Match {
	collection := mongo.Db.Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		log.Fatal(err)
	}

	var results []*MatchModels.Match
	for cur.Next(context.TODO()) {
		var elem MatchModels.Match
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results
}

func GetMatch(id string) *MatchModels.Match {
	var collection = mongo.Db.Collection(COLLECTION)
	result := MatchModels.Match{}
	err := collection.FindOne(context.TODO(), bson.D{{"matchid", id}}).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	return &result
}

func CreateMatch(match MatchModels.Match) string {
	_, err := mongo.Db.Collection(COLLECTION).InsertOne(context.TODO(), match)
	if err != nil {
		log.Fatal(err)
	}
	return match.MatchID
}
