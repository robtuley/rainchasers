package main

//  Catalogue from https://app.rainchasers.com/catalogue.json like:
//  { version: "v138",
//    rivers: [ {
//	    uuid: "14be0011-a293-4e0d-89df-c0216cf9fe5e",
//	    river: "Aeron",
//	    section: "Cilau Aeron to Aberaeron",
//	    km: 9,
//	    grade: {"text":"2/3","value":2.5,"max":null},
//	    putin: {"lat":52.2151638,"lng":-4.197429},
//	    takeout: {"lat":52.2436576,"lng":-4.2644627},
//	    desc: "Good training run ...",
//	    directions: "The takeout is at ..."
//    }, { ... etc ... } ]
//  }

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string   `json:"text"`
	Average float32  `json:"value"`
	Max     *float32 `json:"max"`
}

// Section is a paddleable river section
type Section struct {
	UUID        string  `json:"uuid"`
	SectionName string  `json:"section"`
	RiverName   string  `json:"river"`
	KM          float32 `json:"km"`
	Grade       Grade   `json:"grade"`
	Putin       LatLng  `json:"putin"`
	Takeout     LatLng  `json:"takeout"`
	Description string  `json:"desc"`
	Directions  string  `json:"directions"`
}
