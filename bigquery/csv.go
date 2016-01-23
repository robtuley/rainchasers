package main

import (
	"encoding/csv"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

func bufferAndFlushToCsv(snapC <-chan gauge.Snapshot, projectId string, bucketName string) (<-chan string, <-chan error, error) {
	csvC := make(chan string)
	errC := make(chan error)

	client, err := google.DefaultClient(context.Background(), storage.ScopeReadWrite)
	if err != nil {
		return csvC, errC, err
	}
	ctx := cloud.NewContext(projectId, client)

	gcs, err := storage.NewClient(ctx)
	if err != nil {
		return csvC, errC, err
	}
	defer gcs.Close()

	bucket := gcs.Bucket(bucketName)

	fileName := "v1/" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".csv"
	w := bucket.Object(fileName).NewWriter(ctx)
	w.ContentType = "text/csv"
	csv := csv.NewWriter(w)

	n := 0
	for s := range snapC {
		n += 1
		r := []string{
			s.Url,
			s.StationUrl,
			s.Name,
			s.RiverName,
			strconv.FormatFloat(float64(s.Lat), 'f', -1, 32),
			strconv.FormatFloat(float64(s.Lg), 'f', -1, 32),
			s.Type,
			s.Unit,
			strconv.FormatInt(s.DateTime.Unix(), 10),
			strconv.FormatFloat(float64(s.Value), 'f', -1, 32),
		}
		if err := csv.Write(r); err != nil {
			errC <- err
		}
		if n > 10 {
			break
		}
	}

	csv.Flush()
	if err := csv.Error(); err != nil {
		errC <- err
	}
	if err := w.Close(); err != nil {
		errC <- err
	}

	return csvC, errC, nil
}
