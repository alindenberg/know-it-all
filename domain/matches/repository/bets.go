package matchrepository

import (
	"fmt"
	"log"
	"errors"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	mongo "github.com/alindenberg/know-it-all/database"
	BetModels "github.com/alindenberg/know-it-all/domain/matches/models"
)

var BETS_COLLECTION = "bets"

func GetAllBetsForUser(userId string) ([]*BetModels.Bet, error) {
	collection := mongo.GetDbClient().Collection(BETS_COLLECTION)

	cur, err := collection.Find(context.TODO(), bson.D{{"userid", userId}}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*BetModels.Bet
	for cur.Next(context.TODO()) {
		var elem BetModels.Bet
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetAllBetsForMatch(matchId string) ([]*BetModels.Bet, error) {
	collection := mongo.GetDbClient().Collection(BETS_COLLECTION)

	cur, err := collection.Find(context.TODO(), bson.D{{"matchid", matchId}}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*BetModels.Bet
	for cur.Next(context.TODO()) {
		var elem BetModels.Bet
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

// func GetBet(id string) (*BetModels.Bet, error) {
// 	collection := mongo.Db.Collection(COLLECTION)
// 	result := BetModels.Bet{}
// 	err := collection.FindOne(context.TODO(), bson.D{{"betid", id}}).Decode(&result)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &result, nil
// }

func CreateBet(bet BetModels.Bet, userId string) error {
	collection := mongo.GetDbClient().Collection(BETS_COLLECTION)

	_, err := collection.InsertOne(context.Background(), bet)
	return err
}

func DeleteBet(betId string, userId string) error {
	collection := mongo.GetDbClient().Collection(BETS_COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"betid", betId}})

	if err != nil {
		return err
	}

	if(result.DeletedCount == 0) {
		return errors.New(fmt.Sprintf("Document with id %s was not found", betId))
	}
	return nil
}

func ResolveBet(id string, won bool) error {
	collection := mongo.GetDbClient().Collection(BETS_COLLECTION)

	update := bson.D{
		{"$set", bson.D{
			{"won", won},
			{"isresolved", true},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"betid", id}}, update)

	return err
}
