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
	"strings"
)

var h = sha256.New()

func ConvertJSONToCSV(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var File []CSVFileJSON
	if err := json.NewDecoder(srcFile).Decode(&File); err != nil {
		return err
	}
	outputFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"TEAM NAMES", "Series Number", "Filename", "Name", "Description", "Gender", "Attributes", "UUID", "Hash"}
	if err := writer.Write(header); err != nil {
		return err
	}
	for _, file := range File {
		var csvRow []string
		csvRow = append(csvRow, file.SeriesNumber, file.FileName, file.Name, file.Description, file.Gender, file.Attributes, file.UUID, file.Hash)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func GetAllLines(data [][]string) {
	var JSONFile []CSVFileJSON
	for j, record := range data {
		//omit header line
		if j > 0 {
			var rec CSVFileJSON
			var chip CHIP_0007
			for i, field := range record {
				switch i {
				case 0:
					rec.TeamName = field
					chip.MintingTool = field
				case 1:
					rec.SeriesNumber = field
					chip.SeriesNumber = field
				case 2:
					rec.FileName = field
					chip.Name = field
				case 3:
					rec.Name = field
				case 4:
					rec.Description = field
					chip.Description = field
				case 5:
					rec.Gender = field
				case 6:
					rec.Attributes = field
					// write the logic for the attributes
					//loop through field
					str := strings.Split(field, "; ")
					for _, fd := range str {
						f := strings.Split(fd, ":")
						attr := Attributes{TraitType: f[0], Value: f[1]}
						fmt.Println(attr)
						chip.Attributes = append(chip.Attributes, attr)
					}
					//chip.Attributes = []Attributes{}
				case 7:
					rec.UUID = field
				}
			}
			//CSVFileJSON FORMAT
			// convert the array to json using the MarshalIndent function
			jsData, err := json.MarshalIndent(rec, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			nftfilename := fmt.Sprintf("csv/%s.json", rec.FileName)
			err = os.WriteFile(nftfilename, jsData, 0644)
			if err != nil {
				log.Fatal(err)
			}
			//add rest of CHIP FORMAT data
			chip.Format = "CHIP-0007"
			chip.SensitiveContent = false
			chip.SeriesTotal = 420
			attributes := CollectionAttribute{
				Type:  "description",
				Value: "Rewards for accomplishments during HNGi9.",
			}
			chip.Collection = Collection{
				Name:       "Zuri NFT Tickets for Free Lunch",
				ID:         "b774f676-c1d5-422e-beed-00ef5510c64d",
				Attributes: []CollectionAttribute{attributes},
			}

			//CHIP_0007 FORMAT
			jsonData, err := json.MarshalIndent(chip, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			chipfilename := fmt.Sprintf("chip-0007/%s.json", chip.Name)
			err = os.WriteFile(chipfilename, jsonData, 0644)
			if err != nil {
				log.Fatal(err)
			}
			chi, err := os.Open(chipfilename)
			if err != nil {
				log.Fatal(err)
			}
			defer chi.Close()

			//Hash the chip-0007 format json file
			if _, err := io.Copy(h, chi); err != nil {
				log.Fatal(err)
			}

			rec.Hash = hex.EncodeToString(h.Sum(nil))
			JSONFile = append(JSONFile, rec)
		}
	}
	jsonData, err := json.MarshalIndent(JSONFile, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := ConvertJSONToCSV("output.json", "filename.output.csv"); err != nil {
		log.Fatal(err)
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

	// read csv data using csv.Reader
	csvReader := csv.NewReader(f)
	record, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// grab all the csv data
	GetAllLines(record)
}
