package db

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ijimmywei/covid-grapher/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetRecords(countryName *string, region *string) ([]*model.Record, error)
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

func (db MongoDB) GetRecords(countryName *string, region *string) ([]*model.Record, error) {
	var res *mongo.Cursor
	var err error

	if countryName != nil {
		res, err = db.collection.Find(context.TODO(), db.filterByCountryName(*countryName))
	} else if region != nil {
		res, err = db.collection.Find(context.TODO(), db.filterByRegion(*region))
	} else {
		res, err = db.collection.Find(context.TODO(), bson.M{})
	}

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
		"countriesAndTerritories": ToTitleSnakecase(countryName),
	}
}

func (db MongoDB) filterByRegion(region string) bson.M {
	return bson.M{
		"continentExp": ToTitleSnakecase(region),
	}
}

func ToTitleSnakecase(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Title(str)
	str = strings.Replace(str, " ", "_", -1)

	str = strings.Replace(str, "Of", "of", -1)
	str = strings.Replace(str, "The", "the", -1)
	str = strings.Replace(str, "And", "and", -1)

	return str
}
