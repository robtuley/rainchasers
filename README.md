Rainchasers Gauge Service
=========================

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub, archiving into Google BigQuery and presenting out an API gauge view service for the [Rainchasers PWA](https://app.rainchasers.com).

[![Circle CI](https://circleci.com/gh/rainchasers/com.rainchasers.gauge.svg?style=svg)](https://circleci.com/gh/rainchasers/com.rainchasers.gauge)

[![Code Climate](https://codeclimate.com/github/rainchasers/com.rainchasers.gauge/badges/gpa.svg)](https://codeclimate.com/github/rainchasers/com.rainchasers.gauge)

[![Issue Count](https://codeclimate.com/github/rainchasers/com.rainchasers.gauge/badges/issue_count.svg)](https://codeclimate.com/github/rainchasers/com.rainchasers.gauge)

`/gauge/` Common Library
------------------------

[![GoDoc](https://godoc.org/github.com/rainchasers/com.rainchasers.gauge/gauge?status.png)](https://godoc.org/github.com/rainchasers/com.rainchasers.gauge/gauge)

Library package that covers:

* `Snapshot` type definition
* Avro schema, encoding & decoding
* Google Cloud PubSub publish & consume

`/ea/` Retrieve Latest EA Gauge Values
--------------------------------------

Daemon that polls the [EA Real Time Data API](http://environment.data.gov.uk/flood-monitoring/doc/reference) and publishes a stream of the latest reading values to a PubSub topic. The daemon executes in 3 phases:

1. *Station Discovery* (`discover.go`): paginate through the station summaries to gather the full list of stations.  
2. *Station Metadata* (`detail.go`): for each station get the required meta-data. The need for the Lg/Lat means an individual request needs to be made for each station. 
3. *Measurement Update* (`update.go`): paginates through all measurements and publishes the latest measurement combined (`combine.go`) with the already gathered meta-data.

A standard API http request library (`http.go`) throttles the requests into the EA API. 

In addition to this 3 stage monitoring for the latest readings, after stage (2) the application downloads historical information in batch CSV form (`history.go`) and publishes this to a separate history PubSub channel to ensure that the detailed readings in-between latest updates are captured.

`/bigquery/` Archive Gauge Values to Google Storage & Big Query
---------------------------------------------------------------

ETL daemon that consumes the PubSub stream of both latest and historical gauge values, archives to Google Storage in CSV format and loads into Google Big Query.

Docker > GCE Kubernetes Deployment
----------------------------------

Required applications are built by CircleCI into docker images that are then deployed into a GCE kubernetes cluster using the appropriate `deployment.yml`.