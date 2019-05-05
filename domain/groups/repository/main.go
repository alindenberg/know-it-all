package grouprepository

import (
	"fmt"
	"errors"
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
	collection := mongo.Db.Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), group)
	
	return err
}

func DeleteGroup(id string) error {
	collection := mongo.Db.Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"groupid", id}})

	if err != nil {
		return err
	}

	if(result.DeletedCount == 0) {
		return errors.New(fmt.Sprintf("Document with id %s was not found", id))
	}
	return nil
}