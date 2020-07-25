package db

import (
	"context"
	"fmt"
	"log"

	"github.com/ijimmywei/covid-grapher/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetRecordsByCountryName(countryName string) ([]*model.Record, error)
	GetAllCountries() ([]string, error)
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

func (db MongoDB) GetRecordsByCountryName(countryName string) ([]*model.Record, error) {
	res, err := db.collection.Find(context.TODO(), db.filterByCountryName(countryName))
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

func (db MongoDB) GetAllCountries() ([]string, error) {
	res, err := db.collection.Distinct(context.TODO(), "countriesAndTerritories", bson.D{{}})

	if err != nil {
		log.Printf("Error while fetching countries: %s", err.Error())
		return nil, err
	}
	var p []string
	for _, value := range res {
		s := fmt.Sprint(value)
		p = append(p, s)
	}

	return p, nil
}

func (db MongoDB) filterByCountryName(countryName string) bson.M {
	return bson.M{
		"countriesAndTerritories": countryName,
	}
}
