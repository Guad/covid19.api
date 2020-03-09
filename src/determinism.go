package main

import (
	"sort"
	"strings"
)

func sortRegions(regions []*LastRegionStatus) []*LastRegionStatus {
	sort.SliceStable(regions, func(i, j int) bool {
		c := strings.Compare(regions[i].Country, regions[j].Country)

		if c != 0 {
			return c < 0
		}

		reg := strings.Compare(regions[i].Region, regions[j].Region)
		return reg < 0
	})

	return regions
}

func sortCountries(regions []*LastCountryStatus) []*LastCountryStatus {
	sort.SliceStable(regions, func(i, j int) bool {
		return strings.Compare(regions[i].Name, regions[j].Name) < 0
	})

	q := []*LastCountryStatus{}

	for _, r := range regions {
		q = append(q, r)
	}

	for len(q) > 0 {
		deq := q[0]
		q = q[1:]

		if len(deq.Regions) > 0 {
			sort.SliceStable(deq.Regions, func(i, j int) bool {
				return strings.Compare(deq.Regions[i].Name, deq.Regions[j].Name) < 0
			})

			for _, r := range deq.Regions {
				q = append(q, r)
			}
		}
	}

	return regions
}
