package river

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `firestore:"lat" yaml:"lat"`
	Lng float32 `firestore:"lng" yaml:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string  `firestore:"text" yaml:"text"`
	Average float32 `firestore:"value" yaml:"value"`
	Max     float32 `firestore:"max,omitempty" yaml:"max,omitempty"`
}

// Section is a paddleable river section
type Section struct {
	UUID        string  `firestore:"uuid" yaml:"uuid"`
	Slug        string  `firestore:"slug"` // derived from filename not yaml
	SectionName string  `firestore:"section" yaml:"section"`
	RiverName   string  `firestore:"river" yaml:"river"`
	KM          float32 `firestore:"km" yaml:"km"`
	Grade       Grade   `firestore:"grade" yaml:"grade"`
	Putin       LatLng  `firestore:"putin" yaml:"putin"`
	Takeout     LatLng  `firestore:"takeout" yaml:"takeout"`
	Description string  `firestore:"desc" yaml:"desc"`
	Directions  string  `firestore:"directions" yaml:"directions"`
}
