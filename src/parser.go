package main

import "strings"

func reduceCountries(regions []*LastRegionStatus) []*LastCountryStatus {
	countries := map[string]*LastCountryStatus{}

	for _, row := range regions {
		country := strings.TrimSpace(row.Country)
		region := strings.TrimSpace(row.Region)

		if _, ok := countries[country]; !ok {
			countries[country] = &LastCountryStatus{
				Name:           country,
				ConfirmedCases: 0,
				Deaths:         0,
				Recovered:      0,
				Regions:        []*LastCountryStatus{},

				HistoricCases:     make([]int, 5),
				HistoricDeaths:    make([]int, 5),
				HistoricRecovered: make([]int, 5),
			}
		}

		countries[country].ConfirmedCases += row.ConfirmedCases
		countries[country].Deaths += row.Deaths
		countries[country].Recovered += row.Recovered

		for i := 0; i < 5; i++ {
			countries[country].HistoricCases[i] += row.HistoricCases[i]
			countries[country].HistoricDeaths[i] += row.HistoricDeaths[i]
			countries[country].HistoricRecovered[i] += row.HistoricRecovered[i]
		}

		if region != "" && region != country {
			countries[country].Regions = append(countries[country].Regions,
				&LastCountryStatus{
					Name:           region,
					ConfirmedCases: row.ConfirmedCases,
					Deaths:         row.Deaths,
					Recovered:      row.Recovered,

					HistoricCases:     row.HistoricCases,
					HistoricDeaths:    row.HistoricDeaths,
					HistoricRecovered: row.HistoricRecovered,
				},
			)
		}
	}

	slice := make([]*LastCountryStatus, len(countries))

	i := 0
	for _, v := range countries {
		slice[i] = v
		i++
	}

	return slice
}

func reduceRegions(countries []*LastCountryStatus) []*LastCountryStatus {
	for _, country := range countries {
		subregions := map[string]*LastCountryStatus{}

		for _, region := range country.Regions {
			subr := strings.Split(region.Name, ", ")

			if len(subr) > 1 {
				parent := strings.TrimSpace(subr[1])
				child := strings.TrimSpace(subr[0])

				if _, ok := subregions[parent]; !ok {
					subregions[parent] = &LastCountryStatus{
						Name:           parent,
						ConfirmedCases: 0,
						Deaths:         0,
						Recovered:      0,
						Regions:        []*LastCountryStatus{},

						HistoricCases:     make([]int, 5),
						HistoricDeaths:    make([]int, 5),
						HistoricRecovered: make([]int, 5),
					}
				}

				subregions[parent].ConfirmedCases += region.ConfirmedCases
				subregions[parent].Deaths += region.Deaths
				subregions[parent].Recovered += region.Recovered

				for i := 0; i < 5; i++ {
					subregions[parent].HistoricCases[i] += region.HistoricCases[i]
					subregions[parent].HistoricDeaths[i] += region.HistoricDeaths[i]
					subregions[parent].HistoricRecovered[i] += region.HistoricRecovered[i]
				}

				subregions[parent].Regions = append(subregions[parent].Regions,
					&LastCountryStatus{
						Name:           child,
						ConfirmedCases: region.ConfirmedCases,
						Deaths:         region.Deaths,
						Recovered:      region.Recovered,

						HistoricCases:     region.HistoricCases,
						HistoricDeaths:    region.HistoricDeaths,
						HistoricRecovered: region.HistoricRecovered,
					},
				)
			}
		}

		if len(subregions) > 0 {
			country.Regions = []*LastCountryStatus{}

			for _, subr := range subregions {
				country.Regions = append(country.Regions, subr)
			}
		}
	}

	return countries
}
