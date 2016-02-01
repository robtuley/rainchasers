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
	Id                string
	Bucket            string
	Object            string
	BatchSize         int
	ListenNanoseconds int64
	WriteNanoseconds  int64
}

func csvEncodeAndWrite(projectId string, bucketName string, maxBatchSize int, snapC <-chan gauge.Snapshot) (<-chan CSVFile, <-chan error) {
	csvC := make(chan CSVFile)
	errC := make(chan error)

	// maintain a single go-routine encoding snapC
	// to csv, then flush to google storage

	nextBatchC := make(chan bool)
	go func() {
		for {
			go singleBatchCsvEncodeAndWrite(projectId, bucketName, snapC, csvC, errC, nextBatchC, maxBatchSize)
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
	maxBatchSize int,
) {
	startListenTime := time.Now()
	gContext, gClient, err := getAuthedClient(projectId)
	if err != nil {
		errC <- err
		return
	}
	defer gClient.Close()

	id := strconv.FormatInt(startListenTime.UnixNano(), 10)
	objectName := startListenTime.Format("2006/01/02/") + id + ".csv"

	bucket := gClient.Bucket(bucketName)
	gw := bucket.Object(objectName).NewWriter(gContext)
	gw.ContentType = "text/csv"

	cw := csv.NewWriter(gw)
	nBatch := 0
	timeout := time.Tick(time.Minute * 20)

ThisBatch:
	for {
		select {
		case s := <-snapC:
			nBatch += 1
			r := snap2Record(s)
			if err := cw.Write(r); err != nil {
				errC <- err
			}
			if nBatch >= maxBatchSize {
				nextBatchC <- true
				break ThisBatch
			}
		case <-timeout:
			// write early to maintain data flow and prevent
			// Google auth timeout errors
			nextBatchC <- true
			break ThisBatch
		}
	}

	listenDuration := time.Now().Sub(startListenTime)
	startWriteTime := time.Now()

	cw.Flush()
	if err := cw.Error(); err != nil {
		errC <- err
	}
	if err := gw.Close(); err != nil {
		errC <- err
	}
	if nBatch == 0 {
		return // no records to process
	}

	writeDuration := time.Now().Sub(startWriteTime)
	csvC <- CSVFile{
		Id:                id,
		Bucket:            bucketName,
		Object:            objectName,
		BatchSize:         nBatch,
		ListenNanoseconds: listenDuration.Nanoseconds(),
		WriteNanoseconds:  writeDuration.Nanoseconds(),
	}
}

func getAuthedClient(projectId string) (context.Context, *storage.Client, error) {
	client, err := google.DefaultClient(context.Background(), storage.ScopeReadWrite)
	if err != nil {
		return nil, nil, err
	}
	ctx := cloud.NewContext(projectId, client)

	sClient, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ctx, sClient, nil
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
