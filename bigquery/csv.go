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

type CSVFile struct {
	Id     string
	Bucket string
	Object string
}

func csvEncodeAndWrite(projectId string, bucketName string, batchSize int, snapC <-chan gauge.Snapshot) (<-chan CSVFile, <-chan error, error) {
	csvC := make(chan CSVFile)
	errC := make(chan error)

	// auth with Google storage, and get bucket handle

	client, err := google.DefaultClient(context.Background(), storage.ScopeReadWrite)
	if err != nil {
		return csvC, errC, err
	}
	ctx := cloud.NewContext(projectId, client)

	sClient, err := storage.NewClient(ctx)
	if err != nil {
		return csvC, errC, err
	}
	defer sClient.Close()

	// maintain a single go-routine encoding snapC
	// to csv, then flush to google storage

	nextBatchSignalC := make(chan bool)
	go func() {
		for {
			go singleBatchCsvEncodeAndWrite(ctx, sClient, bucketName, snapC, csvC, errC, nextBatchSignalC, batchSize)
			<-nextBatchSignalC
		}
	}()

	return csvC, errC, nil
}

func singleBatchCsvEncodeAndWrite(
	gContext context.Context,
	gClient *storage.Client,
	bucketName string,
	snapC <-chan gauge.Snapshot,
	csvC chan<- CSVFile,
	errC chan<- error,
	nextBatchSignalC chan<- bool,
	batchSize int,
) {
	now := time.Now()
	id := strconv.FormatInt(now.UnixNano(), 10)
	file := CSVFile{
		Id:     id,
		Bucket: bucketName,
		Object: now.Format("2006/01/02/") + id + ".csv",
	}

	bucket := gClient.Bucket(bucketName)
	gw := bucket.Object(file.Object).NewWriter(gContext)
	gw.ContentType = "text/csv"

	cw := csv.NewWriter(gw)
	n := 0

ThisBatch:
	for s := range snapC {
		n += 1
		r := snap2Record(s)
		if err := cw.Write(r); err != nil {
			errC <- err
		}
		if n >= batchSize {
			nextBatchSignalC <- true
			break ThisBatch
		}
	}

	cw.Flush()

	if err := cw.Error(); err != nil {
		errC <- err
	}
	if err := gw.Close(); err != nil {
		errC <- err
	}

	csvC <- file
}

func snap2Record(s gauge.Snapshot) []string {
	return []string{
		s.InsertId(),
		s.MetricId(),
		s.Url,
		strconv.FormatFloat(float64(s.Lat), 'f', -1, 32),
		strconv.FormatFloat(float64(s.Lg), 'f', -1, 32),
		s.Type,
		s.Unit,
		strconv.FormatInt(s.DateTime.Unix(), 10),
		strconv.FormatFloat(float64(s.Value), 'f', -1, 32),
	}
}
