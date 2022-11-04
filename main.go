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

	header := []string{"Series Number", "Filename", "Name", "Description", "Gender", "Attributes", "UUID", "Hash"}
	if err := writer.Write(header); err != nil {
		return err
	}
	for _, file := range File {
		var csvRow []string
		csvRow = append(csvRow, file.SeriesNumber, file.Filename, file.Name, file.Description, file.Gender, file.Attributes, file.UUID, file.Hash)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func GetAllLines(data [][]string) {
	var JSONFile []CSVFileJSON
	var chipFormat []CHIP_0007
	for j, record := range data {
		//omit header line
		if j > 0 {
			var rec CSVFileJSON
			var chip CHIP_0007
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
			defer fn.Close()
			h := sha256.New()
			if _, err := io.Copy(h, fn); err != nil {
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
	if err := ConvertJSONToCSV("output.json", "output.csv"); err != nil {
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
