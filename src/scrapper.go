package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func buildTimeSeriesURL(name string) string {
	url := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_19-covid-%v.csv"

	return fmt.Sprintf(url, name)
}

func getLastStatus() ([]*LastRegionStatus, string) {
	respch := make(chan *http.Response, 3)

	fetch := func(c chan *http.Response, url string) {
		resp, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		c <- resp
	}

	go fetch(respch, buildTimeSeriesURL("Confirmed"))
	go fetch(respch, buildTimeSeriesURL("Recovered"))
	go fetch(respch, buildTimeSeriesURL("Deaths"))

	stats := map[string]*LastRegionStatus{}

	dataDate := ""

	for i := 0; i < 3; i++ {
		resp := <-respch
		r := csv.NewReader(resp.Body)

		urlpath := resp.Request.URL.Path

		// Skip header
		header, err := r.Read()

		if err != nil {
			panic(err)
		}

		if dataDate == "" {
			dataDate = header[len(header)-1]
		}

		for {
			record, err := r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			key := record[0] + "$" + record[1]

			data, _ := strconv.Atoi(record[len(record)-1])

			if _, ok := stats[key]; !ok {
				lat, _ := strconv.ParseFloat(record[2], 64)
				long, _ := strconv.ParseFloat(record[3], 64)

				stats[key] = &LastRegionStatus{
					Region:  record[0],
					Country: record[1],
					Lat:     lat,
					Long:    long,
				}
			}

			if strings.Contains(urlpath, "Confirmed") {
				stats[key].ConfirmedCases = data
			} else if strings.Contains(urlpath, "Recovered") {
				stats[key].Recovered = data
			} else {
				stats[key].Deaths = data
			}
		}

		resp.Body.Close()
	}

	slice := make([]*LastRegionStatus, len(stats))

	i := 0
	for _, v := range stats {
		slice[i] = v
		i++
	}

	return slice, dataDate
}
