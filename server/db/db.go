package db

import (
	"context"
	"log"

	"github.com/ijimmywei/covid-grapher/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetRecordsByCountryCode(countryCode string) ([]*model.Record, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *MongoDB {
	covid := client.Database("covid").Collection("covid")
	return &MongoDB{
		collection: covid,
	}
}

func (db MongoDB) GetRecordsByCountryCode(countryCode string) ([]*model.Record, error) {
	res, err := db.collection.Find(context.TODO(), db.filterByCountryCode(countryCode))
	if err != nil {
		log.Printf("Error while fetching records: %s", err.Error())
		return nil, err
	}
	var p []*model.Record
	err = res.All(context.TODO(), &p)
	if err != nil {
		log.Printf("Error while decoding records: %s", err.Error())
		return nil, err
	}

	return p, nil
}

func (db MongoDB) filterByCountryCode(countryCode string) bson.M {
	return bson.M{
		"countryterritoryCode": countryCode,
	}
}
