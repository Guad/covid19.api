package main

import "time"

type LastRegionStatus struct {
	Region         string  `json:"region,omitempty"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Long           float64 `json:"long"`
	ConfirmedCases int     `json:"confirmed_cases"`
	Deaths         int     `json:"deaths"`
	Recovered      int     `json:"recovered"`
}

type LastCountryStatus struct {
	Name           string              `json:"name"`
	ConfirmedCases int                 `json:"confirmed_cases"`
	Deaths         int                 `json:"deaths"`
	Recovered      int                 `json:"recovered"`
	Regions        []LastCountryStatus `json:"regions,omitempty"`
}

type FetchData struct {
	Timestamp time.Time   `json:"timestamp,omitempty"`
	DataDate  string      `json:"data_date,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
