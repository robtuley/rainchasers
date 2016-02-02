package main

import (
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud/bigquery"
)

type BigQueryJobStatus struct {
	Id          string
	Label       string
	Nanoseconds int64
}

type BatchStatus struct {
	File  CSVFile
	Jobs  []BigQueryJobStatus
	Error error
}

func loadCSVIntoBigQuery(projectId string, datasetId string, tableId string, csvC <-chan CSVFile) <-chan BatchStatus {
	batchStatusC := make(chan BatchStatus)
	table := &bigquery.Table{
		ProjectID: projectId,
		DatasetID: datasetId,
		TableID:   tableId,
	}

	go func() {
		for file := range csvC {
			go loadSingleCSVFileIntoBigQuery(table, file, batchStatusC)
		}
	}()

	return batchStatusC
}

func loadSingleCSVFileIntoBigQuery(table *bigquery.Table, file CSVFile, batchStatusC chan<- BatchStatus) {
	status := BatchStatus{
		File:  file,
		Jobs:  make([]BigQueryJobStatus, 2),
		Error: nil,
	}

	client, err := bigQueryClient(table.ProjectID)
	if err != nil {
		status.Error = err
		goto endPipeline
	}

	status.Jobs[0], err = copyCSVFileIntoTable(
		client, table,
		file.Bucket, file.Object,
	)
	if err != nil {
		status.Error = err
		goto endPipeline
	}

	status.Jobs[1], err = dedupThisTable(
		client, table,
	)
	if err != nil {
		status.Error = err
		goto endPipeline
	}

endPipeline:
	batchStatusC <- status
}

func dedupThisTable(client *bigquery.Client, table *bigquery.Table) (BigQueryJobStatus, error) {
	sql := `SELECT insertId,metricId,url, lat,lg,type,unit,timestamp,value
            FROM (
              SELECT *, ROW_NUMBER()
              OVER (PARTITION BY insertId)
              row_number,
              FROM ` + table.FullyQualifiedName() + `
            ) WHERE row_number = 1`

	query := &bigquery.Query{
		Q:                sql,
		DefaultProjectID: table.ProjectID,
		DefaultDatasetID: table.DatasetID,
	}

	job, err := client.Copy(
		context.Background(), table,
		query, bigquery.WriteTruncate,
	)

	status := BigQueryJobStatus{
		Id:    job.ID(),
		Label: "dedup." + table.TableID,
	}
	if err != nil {
		return status, err
	}

	wait, err := waitForJobCompletion(job)
	status.Nanoseconds = wait.Nanoseconds()
	if err != nil {
		return status, err
	}

	return status, nil
}

func copyCSVFileIntoTable(client *bigquery.Client, table *bigquery.Table, bucketName string, objectName string) (BigQueryJobStatus, error) {

	gcs := client.NewGCSReference("gs://" + bucketName + "/" + objectName)

	job, err := client.Copy(
		context.Background(), table, gcs,
		bigquery.CreateDisposition(bigquery.CreateIfNeeded),
		bigquery.DestinationSchema(snapshotSchema()),
		bigquery.MaxBadRecords(1),
	)
	status := BigQueryJobStatus{
		Id:    job.ID(),
		Label: "csv_into." + table.TableID,
	}
	if err != nil {
		return status, err
	}

	wait, err := waitForJobCompletion(job)
	status.Nanoseconds = wait.Nanoseconds()
	if err != nil {
		return status, err
	}

	return status, nil
}

func waitForJobCompletion(job *bigquery.Job) (time.Duration, error) {
	startTime := time.Now()
	for range time.Tick(time.Second * 5) {
		s, err := job.Status(context.Background())
		if err != nil {
			return time.Now().Sub(startTime), err
		}
		if !s.Done() {
			continue
		}
		if err := s.Err(); err != nil {
			return time.Now().Sub(startTime), err
		}
		break
	}
	return time.Now().Sub(startTime), nil
}

func snapshotSchema() bigquery.Schema {
	return bigquery.Schema{
		&bigquery.FieldSchema{
			Name:     "insertId",
			Required: true,
			Type:     bigquery.StringFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "metricId",
			Required: true,
			Type:     bigquery.StringFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "url",
			Required: true,
			Type:     bigquery.StringFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "lat",
			Required: false,
			Type:     bigquery.FloatFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "lg",
			Required: false,
			Type:     bigquery.FloatFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "type",
			Required: true,
			Type:     bigquery.StringFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "unit",
			Required: false,
			Type:     bigquery.StringFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "timestamp",
			Required: true,
			Type:     bigquery.TimestampFieldType,
		},
		&bigquery.FieldSchema{
			Name:     "value",
			Required: true,
			Type:     bigquery.FloatFieldType,
		},
	}
}

func bigQueryClient(projectId string) (*bigquery.Client, error) {
	client, err := google.DefaultClient(context.Background(), bigquery.Scope)
	if err != nil {
		return nil, err
	}
	return bigquery.NewClient(client, projectId)
}
