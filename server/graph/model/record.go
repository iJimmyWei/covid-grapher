package model

type Record struct {
	ID                       string `json:"id" bson:"_id"`
	DateRep                  string `json:"dateRep"`
	Cases                    int    `json:"cases"`
	Deaths                   int    `json:"deaths"`
	CountriesAndTerritories  string `json:"countriesAndTerritories"`
	Cumulative_14d_per_10000 string `json:"cumulative_14d_per_10000" bson:"Cumulative_number_for_14_days_of_COVID-19_cases_per_100000"`
}
