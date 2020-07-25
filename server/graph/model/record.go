package model

type Record struct {
	ID                      string `json:"id" bson:"_id"`
	DateRep                 string `json:"dateRep"`
	Cases                   int    `json:"cases"`
	Deaths                  int    `json:"deaths"`
	CountriesAndTerritories string `json:"countriesAndTerritories"`
}
