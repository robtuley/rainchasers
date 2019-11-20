package main

import (
	"strconv"

	"github.com/rainchasers/content"
	"github.com/rainchasers/content/internal/river"
)

func main() {
	r := &result{}
	for _, s := range content.Sections {
		cals, _ := content.Calibrations[s.UUID]
		r.AddToStats(s, cals)
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
