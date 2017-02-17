package main

import (
	"encoding/csv"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"golang.org/x/net/context"
)

type CSVFile struct {
	Id                string
	Bucket            string
	Object            string
	BatchSize         int
	ListenNanoseconds int64
	WriteNanoseconds  int64
}

func csvEncodeAndWrite(
	projectId string,
	bucketName string,
	minBatchSize int,
	minBatchSeconds int,
	snapC <-chan gauge.Snapshot) (<-chan CSVFile, <-chan error) {

	csvC := make(chan CSVFile)
	errC := make(chan error)

	// maintain a single go-routine encoding snapC
	// to csv, then flush to google storage

	nextBatchC := make(chan bool)
	go func() {
		for {
			go singleBatchCsvEncodeAndWrite(projectId, bucketName, snapC, csvC, errC, nextBatchC, minBatchSize, minBatchSeconds)
			<-nextBatchC
		}
	}()

	return csvC, errC
}

func singleBatchCsvEncodeAndWrite(
	projectId string,
	bucketName string,
	snapC <-chan gauge.Snapshot,
	csvC chan<- CSVFile,
	errC chan<- error,
	nextBatchC chan<- bool,
	minBatchSize int,
	minBatchSeconds int,
) {
	startListenTime := time.Now()
	minBatchEnd := startListenTime.Add(time.Second * time.Duration(minBatchSeconds))
	cache := make([]gauge.Snapshot, 0, minBatchSize)

	for s := range snapC {
		cache = append(cache, s)
		if minBatchEnd.After(time.Now()) {
			continue
		}
		if len(cache) >= minBatchSize {
			break
		}
	}

	listenDuration := time.Now().Sub(startListenTime)
	nextBatchC <- true

	startWriteTime := time.Now()

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		errC <- err
		return
	}
	defer client.Close()

	id := strconv.FormatInt(startListenTime.UnixNano(), 10)
	objectName := startListenTime.Format("2006/01/02/") + id + ".csv"

	bucket := client.Bucket(bucketName)
	gw := bucket.Object(objectName).NewWriter(ctx)
	gw.ContentType = "text/csv"

	cw := csv.NewWriter(gw)
	for _, s := range cache {
		r := snap2Record(s)
		if err := cw.Write(r); err != nil {
			errC <- err
		}
	}
	cw.Flush()
	if err := cw.Error(); err != nil {
		errC <- err
		return
	}
	if err := gw.Close(); err != nil {
		errC <- err
		return
	}

	writeDuration := time.Now().Sub(startWriteTime)
	csvC <- CSVFile{
		Id:                id,
		Bucket:            bucketName,
		Object:            objectName,
		BatchSize:         len(cache),
		ListenNanoseconds: listenDuration.Nanoseconds(),
		WriteNanoseconds:  writeDuration.Nanoseconds(),
	}
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
