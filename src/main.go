package main

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func cleanDir(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			panic(err)
		}

		err = os.Remove(path)

		return err
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	status, date := getLastStatus()
	countries := reduceCountries(status)
	countries = reduceRegions(countries)

	now := time.Now()

	// Delete older files
	cleanDir("../docs")

	// Write JSON to docs
	encodeToFile("../docs/global.json", FetchData{
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
