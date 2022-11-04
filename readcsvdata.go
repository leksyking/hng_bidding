package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
					if field == "" {
						field = teamName
					}
					teamName = field
					rec.TeamName = teamName
					chip.MintingTool = teamName
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
						chip.Attributes = append(chip.Attributes, attr)
					}
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
