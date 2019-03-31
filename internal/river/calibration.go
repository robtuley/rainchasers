package river

// Calibration is a referenced gauge related to a section
type Calibration struct {
	URL         string
	Description string
	Minimum     map[Level]float32
}

// LevelAt provides the level state at a certain reading
func (c Calibration) LevelAt(value float32) Level {
	if len(c.Minimum) == 0 {
		return Unknown
	}

	state := Empty
	for lvl, min := range c.Minimum {
		if value >= min && lvl > state {
			state = lvl
		}
	}

	return state
}
