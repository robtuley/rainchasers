package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/rainchasers/content/internal/river"
	"github.com/rainchasers/report"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FireWriter handles firestore writes
type FireWriter struct {
	Client       *firestore.Client
	Collection   *firestore.CollectionRef
	AlgoliaIndex algoliasearch.Index
	Timeout      time.Duration
}

// NewFireWriter creates a firestore writer
func NewFireWriter(projectID string, algoliaAppID string, algoliaAPIKey string) (*FireWriter, report.Span) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	span := report.StartSpan("firewriter.connect")
	span = span.Field("project_id", projectID)
	span = span.Field("algolia_app_id", algoliaAppID)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, span.End(err)
	}

	algoliaClient := algoliasearch.NewClient(algoliaAppID, algoliaAPIKey)
	algoliaIndex := algoliaClient.InitIndex("rivers")

	return &FireWriter{
		Client:       client,
		Collection:   client.Collection("rivers"),
		AlgoliaIndex: algoliaIndex,
		Timeout:      10 * time.Second,
	}, span.End()
}

// Close releases any resources
func (fw *FireWriter) Close() {
	if fw.Client != nil {
		fw.Client.Close()
	}
}

// LoadAndUpdate saves any update to firestore
func (fw *FireWriter) LoadAndUpdate(ctx context.Context, s river.Section) (*Record, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, fw.Timeout)
	defer cancel()

	span := report.StartSpan("firestore.loadandupdate").Field("uuid", s.UUID)

	// get existing river data to check against
	record, sp := fw.Load(ctx, s.UUID)
	span = span.Child(sp)
	if err := span.Err(); err != nil {
		return record, span.End()
	}

	// if river has not changed, do nothing
	if record != nil {
		if checksum(s) == checksum(record.Section) {
			return record, span.End()
		}
	}

	// write new data with no snapshot gauge readings or state only if river
	// definition has changed
	new := Record{
		Section: s,
		Level: Level{
			Label:  river.Unknown.String(),
			Reason: "Not yet calibrated against nearby gauges",
		},
		Measures: make([]Measure, 0),
	}
	sp = fw.Store(ctx, &new)
	span = span.Child(sp)
	if err := span.Err(); err != nil {
		return &new, span.End()
	}

	return &new, span.End()
}

// Load retrieves the latest river from firestore
func (fw *FireWriter) Load(ctx context.Context, uuid string) (*Record, report.Span) {
	ctx, cancel := context.WithTimeout(ctx, fw.Timeout)
	defer cancel()

	span := report.StartSpan("firestore.load").Field("uuid", uuid)

	doc, err := fw.Collection.Doc(uuid).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// in this case the document simply does not exist so
			// return a nil pointer to indicate this and no span
			// error
			return nil, span.End()
		}
		return nil, span.End(err)
	}
	var record Record
	if err := doc.DataTo(&record); err != nil {
		// in this case assume that the data has changed format or
		// similar and this is not an error in so much as an indication
		// to re-write.
		span = span.Field("corruption", err.Error())
		return nil, span.End()
	}
	return &record, span.End()
}

// Store saves the river to firestore
func (fw *FireWriter) Store(ctx context.Context, record *Record) report.Span {
	ctx, cancel := context.WithTimeout(ctx, fw.Timeout)
	defer cancel()

	uuid := record.Section.UUID

	// write to firestore
	fireSpan := report.StartSpan("firestore.store").Field("uuid", uuid)
	_, err := fw.Collection.Doc(uuid).Set(ctx, record)
	if err != nil {
		return fireSpan.End(err)
	}
	fireSpan = fireSpan.End()

	// write to algolia
	aSpan := report.StartSpan("algolia.store").Field("uuid", uuid)
	s := record.Section
	l := record.Level
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
	_, err = fw.AlgoliaIndex.UpdateObject(object)
	if err != nil {
		return fireSpan.FollowedBy(aSpan.End(err))
	}
	aSpan = aSpan.End()

	return fireSpan.FollowedBy(aSpan)
}
