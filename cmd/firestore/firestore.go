package main

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FireRiver is the river as written to firestore
type FireRiver struct {
	Section   Section
	State     *FireState
	Snapshots map[string]gauge.Snapshot
}

// FireState is the calculated river current state
type FireState struct {
	EventTime      time.Time
	ProcessedTime  time.Time
	NormalisedRank float32 // e.g. normalised calibration
	HumanLabel     string  // e.g. "high"
	HumanSource    string  // e.g. "0.98 at Hafod Wydr gauge"
}

// FireWriter handles firestore writes
type FireWriter struct {
	Client     *firestore.Client
	Collection *firestore.CollectionRef
	Timeout    time.Duration
}

// NewFireWriter creates a firestore writer
func NewFireWriter(projectID string) (*FireWriter, report.Span) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	span := report.StartSpan("firestore.connect")
	span = span.Field("project_id", projectID)

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, span.End(err)
	}

	return &FireWriter{
		Client:     client,
		Collection: client.Collection("rivers"),
		Timeout:    10 * time.Second,
	}, span.End()
}

// Close releases any resources
func (fw *FireWriter) Close() {
	if fw.Client != nil {
		fw.Client.Close()
	}
}

// UpdateSection saves an updated section to firestore
func (fw *FireWriter) UpdateSection(ctx context.Context, s Section) report.Span {
	ctx, cancel := context.WithTimeout(ctx, fw.Timeout)
	defer cancel()

	span := report.StartSpan("firestore.section").Field("uuid", s.UUID)

	// get existing river data to check against
	river, sp := fw.Load(ctx, s.UUID)
	span = span.Child(sp)
	if err := span.Err(); err != nil {
		return span.End()
	}

	// if river has not changed, do nothing
	if river != nil {
		if s.Checksum() == river.Section.Checksum() {
			return span.End()
		}
	}

	// write new data with no snapshot gauge readings or state only if river
	// definition has changed
	new := FireRiver{
		Section: s,
	}
	sp = fw.Store(ctx, &new)
	span = span.Child(sp)
	if err := span.Err(); err != nil {
		return span.End()
	}

	return span.End()
}

// Load retrieves the latest river from firestore
func (fw *FireWriter) Load(ctx context.Context, uuid string) (*FireRiver, report.Span) {
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
	var river FireRiver
	if err := doc.DataTo(&river); err != nil {
		// in this case assume that the data has changed format or
		// similar and this is not an error in so much as an indication
		// to re-write.
		span = span.Field("corruption", err.Error())
		return nil, span.End()
	}
	return &river, span.End()
}

// Store saves the river to firestore
func (fw *FireWriter) Store(ctx context.Context, river *FireRiver) report.Span {
	ctx, cancel := context.WithTimeout(ctx, fw.Timeout)
	defer cancel()

	uuid := river.Section.UUID
	span := report.StartSpan("firestore.store").Field("uuid", uuid)

	_, err := fw.Collection.Doc(uuid).Set(ctx, river)
	if err != nil {
		return span.End(err)
	}

	return span.End()
}
