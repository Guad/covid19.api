package main

import (
	"encoding/json"
	"os"
	"time"
)

func encodeToFile(path string, data interface{}) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(data)

	if err != nil {
		panic(err)
	}
}

func main() {
	status, date := getLastStatus()
	countries := reduceCountries(status)
	countries = reduceRegions(countries)

	now := time.Now()

	// Write JSON to docs
	encodeToFile("../docs/last.json", FetchData{
		Data:      countries,
		Timestamp: now,
		DataDate:  date,
	})
	encodeToFile("../docs/flat.json", FetchData{
		Data:      status,
		Timestamp: now,
		DataDate:  date,
	})
}
