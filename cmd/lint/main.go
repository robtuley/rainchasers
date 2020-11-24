package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/robtuley/rainchasers"
	"github.com/robtuley/rainchasers/internal/river"
)

func main() {
	r := &result{}

sectionLoop:
	for _, s := range rainchasers.Sections {
		cals, _ := rainchasers.Calibrations[s.UUID]
		r.AddToStats(s, cals)

		if len(cals) == 0 {
			fmt.Print(".")
			c, err := findCandidateForCalibration(s)
			if err != nil {
				log.Fatal(err)
			}
			if c != "" {
				fmt.Println("\nRiver " + s.Slug + " has a candidate gauge:")
				fmt.Println(c + "\n")
				break sectionLoop
			}
			// no candidate found, move on
		}
	}
	r.Print()
}

type result struct {
	Count               int
	CountCalibratedOne  int
	CountCalibratedMany int
}

func (r *result) AddToStats(s river.Section, cals []river.Calibration) {
	r.Count++
	if len(cals) > 1 {
		r.CountCalibratedMany++
	}
	if len(cals) == 1 {
		r.CountCalibratedOne++
	}
}

func (r *result) Print() {
	println("Content Lint")
	println("============")
	println("Total: " + strconv.Itoa(r.Count))
	println("Calibrated with One: " + strconv.Itoa(r.CountCalibratedOne))
	println("Calibrated with Many: " + strconv.Itoa(r.CountCalibratedMany))
}

func findCandidateForCalibration(s river.Section) (string, error) {
	// public search-only algolia app and API keys
	appID := "68SCVTV3KD"
	apiKey := "99fe024e02d9f69215cfc5634e5466dc"
	index := algoliasearch.NewClient(appID, apiKey).InitIndex("stations")

	// search for closest gauge, ideally upstream to the putin
	res, err := index.Search(s.RiverName, algoliasearch.Map{})
	if err != nil {
		return "", err
	}
	for _, hit := range res.Hits {
		return "Gauge " + fmt.Sprintf("%v", hit["objectID"]) + " at " + fmt.Sprintf("%v", hit["human_url"]), nil
	}

	return "", nil
}
