package river

// Calibration is a referenced gauge related to a section
type Calibration struct {
	URL         string `firestore:"data_url"`
	Description string `firestore:"desc"`

	// Minimum is a map of the minimum values for each level
	//
	// Note that because we need to write this map to and from firestore
	// we cannot use the native Level type as the key; instead we convert
	// to/from a string.
	Minimum map[string]float32 `firestore:"minimum"`
}

// LevelAt provides the level state at a certain reading
func (c Calibration) LevelAt(value float32) Level {
	if len(c.Minimum) == 0 {
		return Unknown
	}

	state := Empty
	for strLevel, minValue := range c.Minimum {
		if value >= minValue {
			lvl := StringToLevel(strLevel)
			if lvl > state {
				state = lvl
			}
		}
	}

	return state
}
