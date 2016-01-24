package main

import (
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud/bigquery"
)

func loadCSVIntoBigQuery(projectId string, datasetId string, tableId string, csvC <-chan CSVFile) (<-chan error, error) {
	errC := make(chan error)

	client, err := google.DefaultClient(context.Background(), bigquery.Scope)
	if err != nil {
		return errC, err
	}
	bqClient, err := bigquery.NewClient(client, projectId)
	if err != nil {
		return errC, err
	}

	for f := range csvC {
		go func(file CSVFile) {
			tempTableId := tableId + "_" + file.Id
			err := loadSingleCSVFileIntoBigQuery(
				bqClient, projectId,
				file.Bucket, file.Object,
				datasetId, tempTableId,
			)
			if err != nil {
				errC <- err
			}

		}(f)
	}

	return errC, err
}

func loadSingleCSVFileIntoBigQuery(client *bigquery.Client, projectId string, bucketName string, objectName string, datasetId string, tableId string) error {

	table := bigquery.Table{
		ProjectID: projectId,
		DatasetID: datasetId,
		TableID:   tableId,
	}

	gcs := client.NewGCSReference("gs://" + bucketName + "/" + objectName)

	schema := bigquery.Schema{
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

	job, err := client.Copy(
		context.Background(), &table, gcs,
		bigquery.CreateDisposition(bigquery.CreateIfNeeded),
		bigquery.DestinationSchema(schema),
		bigquery.MaxBadRecords(1),
	)
	if err != nil {
		return err
	}

	for range time.Tick(time.Second * 5) {
		status, err := job.Status(context.Background())
		if err != nil {
			return err
		}
		if !status.Done() {
			continue
		}
		if err := status.Err(); err != nil {
			return err
		}
		break
	}

	return nil
}
