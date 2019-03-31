package river

// Level represents the flow level of a river
type Level int

const (
	// Unknown means level is not known
	Unknown Level = -1
	// Empty means river is too low to paddle
	Empty Level = 0
	// Scrape is just paddleable
	Scrape Level = 1
	// Low river level
	Low Level = 2
	// Medium river level
	Medium Level = 3
	// High river level
	High Level = 4
	// Huge means the river level is unusually high
	Huge Level = 5
	// TooHigh means the river level is too high to paddle
	TooHigh Level = 6
)

func (l Level) String() string {
	switch l {
	case Empty:
		return "empty"
	case Scrape:
		return "scrape"
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Huge:
		return "huge"
	case TooHigh:
		return "too_high"
	}
	return "unknown"
}

// StringToLevel converts a string to a level value
func StringToLevel(str string) Level {
	switch str {
	case Empty.String():
		return Empty
	case Scrape.String():
		return Scrape
	case Low.String():
		return Low
	case Medium.String():
		return Medium
	case High.String():
		return High
	case Huge.String():
		return Huge
	case TooHigh.String():
		return TooHigh
	}
	return Unknown
}
