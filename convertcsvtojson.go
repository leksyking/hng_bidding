package main

import (
	"encoding/csv"
	"encoding/json"
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

	header := []string{"TEAM NAMES", "Series Number", "Filename", "Name", "Description", "Gender", "Attributes", "UUID", "Hash"}
	if err := writer.Write(header); err != nil {
		return err
	}
	for _, file := range File {
		var csvRow []string
		csvRow = append(csvRow, file.TeamName, file.SeriesNumber, file.FileName, file.Name, file.Description, file.Gender, file.Attributes, file.UUID, file.Hash)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}
