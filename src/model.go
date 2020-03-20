package main

type LastRegionStatus struct {
	Region         string  `json:"region,omitempty"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Long           float64 `json:"long"`
	ConfirmedCases int     `json:"confirmed_cases"`
	Deaths         int     `json:"deaths"`
	Recovered      int     `json:"recovered"`

	HistoricCases     []int `json:"historic_cases"`
	HistoricDeaths    []int `json:"historic_deaths"`
	HistoricRecovered []int `json:"historic_recovered"`
}

type LastCountryStatus struct {
	Name           string `json:"name"`
	ConfirmedCases int    `json:"confirmed_cases"`
	Deaths         int    `json:"deaths"`
	Recovered      int    `json:"recovered"`

	HistoricCases     []int `json:"historic_cases"`
	HistoricDeaths    []int `json:"historic_deaths"`
	HistoricRecovered []int `json:"historic_recovered"`

	Regions []*LastCountryStatus `json:"regions,omitempty"`
}

type FetchData struct {
	DataDate string      `json:"data_date,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
