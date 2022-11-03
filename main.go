package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// a struct for storing CSV lines and annotate with JSON struct field tags
type CSVFileJSON struct {
	SeriesNumber int    `json:"series_number"`
	Filename     string `json:"filename"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Gender       string `json:"gender"`
	Attributes   string `json:"attributes"`
	UUID         string `json:"uuid"`
	Hash         string `json:"hash"`
}

func GetLine(data [][]string) {
	var rec CSVFileJSON
	for j, record := range data {
		//omit header line
		if j > 0 {
			for i, field := range record {
				switch i {
				case 0:
					//rec.SeriesNumber =
				case 1:
					rec.Filename = field
				case 2:
					rec.Name = field
				case 3:
					rec.Description = field
				case 4:
					rec.Gender = field
				case 5:
					rec.Attributes = field
				case 6:
					rec.UUID = field
				}

			}
			// convert the array to json using the MarshalIndent function
			jsonData, err := json.MarshalIndent(rec, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			// add the data in filename.json
			fmt.Println(string(jsonData))
		}
	}
}

func main() {
	// open file
	f, err := os.Open("HNGi9.csv")
	if err != nil {
		log.Fatal(err)
	}
	// close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)

	record, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	GetLine(record)
}
