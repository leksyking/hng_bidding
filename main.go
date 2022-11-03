package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// a struct for storing CSV lines and annotate with JSON struct field tags
type CSVFileJSON struct {
	SeriesNumber string `json:"series_number"`
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
		if j > 0 && j < 3 {
			for i, field := range record {
				switch i {
				case 0:
					rec.SeriesNumber = field
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
			nftfilename := fmt.Sprintf("%s.json", rec.Filename)
			err = os.WriteFile(nftfilename, jsonData, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fn, err := os.Open(nftfilename)
			if err != nil {
				log.Fatal(err)
			}
			h := sha256.New()
			if _, err := io.Copy(h, fn); err != nil {
				log.Fatal(err)
			}
			rec.Hash = hex.EncodeToString(h.Sum(nil))
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
