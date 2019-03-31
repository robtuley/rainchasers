package river

import (
	"testing"
)

func TestLevelAt(t *testing.T) {
	c := Calibration{
		Minimum: make(map[string]float32),
	}

	if state := c.LevelAt(1.0); state != Unknown {
		t.Fatal("no readings and has level", state)
	}

	c.Minimum[Scrape.String()] = 1.2
	c.Minimum[High.String()] = 1.7

	expect := map[Level]float32{
		Empty:  1.19,
		Scrape: 1.21,
		High:   1.71,
	}
	for lvl, value := range expect {
		if state := c.LevelAt(value); state != lvl {
			t.Fatal(value, "has level", state)
		}
	}
}
