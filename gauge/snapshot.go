package gauge

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

type Snapshot struct {
	DataURL   string
	HumanURL  string
	Name      string
	RiverName string
	Lat       float32
	Lg        float32
	Type      string
	Unit      string
	DateTime  time.Time
	Value     float32
}

type SnapshotUpdate struct {
	MetricID string
	DateTime time.Time
	Value    float32
}

func (s Snapshot) Apply(u SnapshotUpdate) Snapshot {
	if s.MetricID() != u.MetricID {
		panic("Snapshot update " + u.MetricID + " applied to " + s.MetricID())
	}
	s.DateTime = u.DateTime
	s.Value = u.Value
	return s
}

func (s Snapshot) InsertID() string {
	h := sha256.New()
	io.WriteString(h, s.DateTime.Format(time.RFC822)+strings.ToLower(s.DataURL))
	return hex.EncodeToString(h.Sum(nil))
}

func CalculateMetricID(dataURL string) string {
	h := sha256.New()
	io.WriteString(h, strings.ToLower(dataURL))
	return hex.EncodeToString(h.Sum(nil))
}

func (s Snapshot) MetricID() string {
	return CalculateMetricID(s.DataURL)
}
