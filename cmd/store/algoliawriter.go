package main

import (
	"context"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/rainchasers/report"
)

// AlgoliaWriter handles algolia writes
type AlgoliaWriter struct {
	RiverIndex   algoliasearch.Index
	MeasureIndex algoliasearch.Index
	Timeout      time.Duration
}

// NewAlgoliaWriter creates a algolia writer
func NewAlgoliaWriter(appID string, APIKey string) *AlgoliaWriter {
	client := algoliasearch.NewClient(appID, APIKey)
	return &AlgoliaWriter{
		RiverIndex:   client.InitIndex("rivers"),
		MeasureIndex: client.InitIndex("measures"),
		Timeout:      10 * time.Second,
	}
}

// StoreRecord saves a river record to algolia
func (aw *AlgoliaWriter) StoreRecord(ctx context.Context, record *Record) report.Span {
	ctx, cancel := context.WithTimeout(ctx, aw.Timeout)
	defer cancel()

	uuid := record.Section.UUID
	s := record.Section
	l := record.Level

	span := report.StartSpan("algolia.store").Field("uuid", s.UUID)
	object := algoliasearch.Object{
		"objectID":      uuid,
		"slug":          s.Slug,
		"section":       s.SectionName,
		"river":         s.RiverName,
		"grade":         s.Grade.Human,
		"grade_numeric": s.Grade.Average,
		"desc":          s.Description,
		"km":            s.KM,
		"_geoloc": map[string]float32{
			"lat": s.Putin.Lat,
			"lng": s.Putin.Lng,
		},
		"level_label":     l.Label,
		"level_reason":    l.Reason,
		"level_timestamp": l.EventTime,
	}
	_, err := aw.RiverIndex.UpdateObject(object)
	if err != nil {
		return span.End(err)
	}

	return span.End()
}
