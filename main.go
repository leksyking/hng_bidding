package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// a struct for storing CSV lines and annotate with JSON struct field tags
type CSVFileJSON struct {
	SeriesNumber string `series_number`
	Filename     string `json:"filename"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Gender       string `json:"gender"`
	Attributes   string `json:"attributes"`
	UUID         string `json:"uuid"`
	Hash         string `json:"hash"`
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
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// convert the line to json
		var rec CSVFileJSON
		for i, field := range record {
			switch i {
			case 0:
				rec.SeriesNumber = field
			case 1:
				rec.Filename = field
			case 2:
				rec.Name = field
			case 4:
				rec.Description = field
			case 5:
				rec.Gender = field
			case 6:
				rec.Attributes = field
			case 7:
				rec.UUID = field
			}
		}
		jsonData, err := json.MarshalIndent(rec, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonData))
	}
}
