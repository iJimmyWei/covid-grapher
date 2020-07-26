package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileUrl := "https://opendata.ecdc.europa.eu/covid19/casedistribution/json"
	path := "mongo.init"
	collectionName := "covid"

	err := DownloadFile(path, fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)

	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Parse to allow mongo to insert these records in our collection
	newContents := strings.Replace(string(f), "\"records\" : ", "db."+collectionName+".insert(", -1)
	newContents = newContents[1:]
	newContents = newContents[:len(newContents)-1]
	newContents = newContents + ");"

	// Create indexes
	newContents = newContents + "db." + collectionName + ".createIndex( { countriesAndTerritories: 1 } );"

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}

// Write as we download, no need to hold it in memory
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
