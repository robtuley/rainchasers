package main

import (
	"strconv"

	"github.com/rainchasers/content"
)

func main() {
	r := result{}
	for _, s := range content.Sections {
		r.Count++
		measures, isCalibrated := content.Calibrations[s.UUID]
		if isCalibrated {
			if len(measures) > 1 {
				r.CountCalibratedMany++
			}
			if len(measures) == 1 {
				r.CountCalibratedOne++
			}

		}
	}
	r.Print()
}

type result struct {
	Count               int
	CountCalibratedOne  int
	CountCalibratedMany int
}

func (r result) Print() {
	println("Content Lint")
	println("============")
	println("Total: " + strconv.Itoa(r.Count))
	println("Calibrated with One: " + strconv.Itoa(r.CountCalibratedOne))
	println("Calibrated with Many: " + strconv.Itoa(r.CountCalibratedMany))
}
