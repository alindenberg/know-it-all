package grouprepository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	mongo "github.com/alindenberg/know-it-all/database"
	GroupModels "github.com/alindenberg/know-it-all/domain/groups/models"
)

var COLLECTION = "groups"

func GetAllGroups() ([]*GroupModels.Group, error) {
	collection := mongo.Db.Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*GroupModels.Group
	for cur.Next(context.TODO()) {
		var elem GroupModels.Group
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetGroup(id string) (*GroupModels.Group, error) {
	collection := mongo.Db.Collection(COLLECTION)
	result := GroupModels.Group{}
	err := collection.FindOne(context.TODO(), bson.D{{"groupid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateGroup(group GroupModels.Group) error {
	_, err := mongo.Db.Collection(COLLECTION).InsertOne(context.TODO(), group)
	if err != nil {
		return err
	}
	return nil
}