package main

// a struct for storing CSV lines and annotate with JSON struct field tags
type CSVFileJSON struct {
	TeamName     string `json:"team_name"`
	SeriesNumber string `json:"series_number"`
	FileName     string `json:"filename"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Gender       string `json:"gender"`
	Attributes   string `json:"attributes"`
	UUID         string `json:"uuid"`
	Hash         string `json:"hash"`
}

type CHIP_0007 struct {
	Format           string       `json:"format"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	MintingTool      string       `json:"minting_tool"`
	SensitiveContent bool         `json:"sensitive_content"`
	SeriesNumber     string       `json:"series_number"`
	SeriesTotal      uint         `json:"series_total"`
	Attributes       []Attributes `json:"attributes"`
	Collection       Collection   `json:"collection"`
}

type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Collection struct {
	Name       string                `json:"name"`
	ID         string                `json:"id"`
	Attributes []CollectionAttribute `json:"attributes"`
}

type CollectionAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
