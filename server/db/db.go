package db

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/ijimmywei/covid-grapher/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetRecords(countryName *string) ([]*model.Record, error)
	GetRecordsByRegion(region string) ([]*model.Region, error)
	GetAllCountries() ([]string, error)
	GetAllRegions() ([]string, error)
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

func (db MongoDB) GetRecords(countryName *string) ([]*model.Record, error) {
	var res *mongo.Cursor
	var err error

	if countryName != nil {
		res, err = db.collection.Find(context.TODO(), db.filterByCountryName(*countryName))
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

func (db MongoDB) GetRecordsByRegion(region string) ([]*model.Region, error) {
	var res *mongo.Cursor
	var err error

	res, err = db.collection.Find(context.TODO(), db.filterByRegion(region))

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

	// Region searches compress all countries' data
	m := make(map[string]*model.Region)

	for _, record := range p {
		currentCases := 0
		currentDeaths := 0
		if m[record.DateRep] != nil {
			currentCases = m[record.DateRep].Cases
			currentDeaths = m[record.DateRep].Deaths
		}

		newCases := currentCases + record.Cases
		newDeaths := currentDeaths + record.Deaths

		m[record.DateRep] = &model.Region{
			DateRep: record.DateRep,
			Cases:   newCases,
			Deaths:  newDeaths,
		}
	}

	var finalData []*model.Region
	for _, allRecords := range m {
		finalData = append(finalData, allRecords)
	}

	sort.Sort(SortBy(finalData))

	return finalData, nil
}

type SortBy []*model.Region

var layout = "02/01/2006"

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	t1, _ := time.Parse(layout, a[i].DateRep)
	t2, _ := time.Parse(layout, a[j].DateRep)

	return t1.Before(t2)
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

func (db MongoDB) GetAllRegions() ([]string, error) {
	res, err := db.collection.Distinct(context.TODO(), "continentExp", bson.D{{}})

	if err != nil {
		log.Printf("Error while fetching regions: %s", err.Error())
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
