package main

import (
	"crypto/sha256"
	"encoding/csv"
	"log"
	"os"
)

var (
	h        = sha256.New() //sha256 hash
	teamName string         //team name global variable
)

func main() {
	// open csv file
	f, err := os.Open("HNGi9.csv")
	if err != nil {
		log.Fatal(err)
	}
	// close the file at the end of the program
	defer f.Close()

	// read csv data using csv.Reader
	csvReader := csv.NewReader(f)
	record, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// grab all the csv data
	GetAllLines(record)
}
