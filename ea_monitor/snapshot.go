package main

import (
	"time"
)

type Snapshot struct {
	Url        string
	StationUrl string
	Name       string
	RiverName  string
	Lat        float32
	Lg         float32
	Type       string
	Unit       string
	DateTime   time.Time
	Value      float32
}

type SnapshotUpdate struct {
	Url      string
	DateTime time.Time
	Value    float32
}

func (s Snapshot) Apply(u SnapshotUpdate) Snapshot {
	if s.Url != u.Url {
		panic("Snapshot update " + u.Url + " applied to " + s.Url)
	}
	s.DateTime = u.DateTime
	s.Value = u.Value
	return s
}
