package betrepository

import (
	"fmt"
	"log"
	"strings"
	"errors"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	mongo "github.com/alindenberg/know-it-all/database"
	BetModels "github.com/alindenberg/know-it-all/domain/bets/models"
)

var COLLECTION = "userbets"

func GetAllBetsForUser(userId string) ([]*BetModels.Bet, error) {
	collection := mongo.Db.Collection(COLLECTION)

	userBet := BetModels.UserBets{}
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userId}}).Decode(&userBet)
	if err != nil {
		return nil, err
	}

	return userBet.Bets, nil
}

func GetAllBetsForMatch(matchId string) ([]*BetModels.Bet, error) {
	collection := mongo.Db.Collection(COLLECTION)

	userBet := BetModels.UserBets{}
	err := collection.FindOne(context.TODO(), bson.D{{"matchid", matchId}}).Decode(&userBet)
	if err != nil {
		return nil, err
	}

	return userBet.Bets, nil
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
	collection := mongo.Db.Collection(COLLECTION)

	userBet := BetModels.UserBets{}
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userId}}).Decode(&userBet)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			// Logic for no document found
			userBet = BetModels.UserBets{
				UserID: userId,
				Bets: []*BetModels.Bet {
					&bet,
				},
			}

			_, err = mongo.Db.Collection(COLLECTION).InsertOne(context.TODO(), userBet)
			return err
		}
		return err
	}

	update := bson.D{
		{"$push", bson.D{
			{"bets", bet},
		}},
	}
	_, err = mongo.Db.Collection(COLLECTION).UpdateOne(context.TODO(), bson.D{{"userid", userId}}, update)

	return err

}

func DeleteBet(id string, userId string) error {
	collection := mongo.Db.Collection(COLLECTION)

	userBet := BetModels.UserBets{}
	err := collection.FindOne(context.TODO(), bson.D{{"userid", userId}}).Decode(&userBet)
	if err != nil {
		return err
	}

	var betToBeDeleted *BetModels.Bet

	for _, bet := range userBet.Bets {
      if bet.BetID == id {
          betToBeDeleted = bet
      }
  }

	if betToBeDeleted != nil {
		update := bson.D{
			{"$pull", bson.D{
				{"bets", betToBeDeleted},
			}},
		}
		_, err = mongo.Db.Collection(COLLECTION).UpdateOne(context.TODO(), bson.D{{"userid", userId}}, update)

		return err
	}

	return errors.New(fmt.Sprintf("Bet %s not found", id))
}

func ResolveBet(id string, won bool) error {
	log.Println("repo resolving bet")
	update := bson.D{
		{"$set", bson.D{
			{"won", won},
			{"isresolved", true},
		}},
	}
	_, err := mongo.Db.Collection(COLLECTION).UpdateOne(context.TODO(), bson.D{{"betid", id}}, update)

	return err
}
