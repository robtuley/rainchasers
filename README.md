Rainchasers Gauge Service
=========================

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub, archiving into Google BigQuery and presenting out a HTML gauge view service for the Rainchasers site.

[![Circle CI](https://circleci.com/gh/rainchasers/com.rainchasers.gauge.svg?style=svg)](https://circleci.com/gh/rainchasers/com.rainchasers.gauge)

`/gauge/` Common Library
------------------------

[![GoDoc](https://godoc.org/github.com/rainchasers/com.rainchasers.gauge/gauge?status.png)](https://godoc.org/github.com/rainchasers/com.rainchasers.gauge/gauge)

Library package that covers:

* `Snapshot` type definition
* Avro schema, encoding & decoding
* Google Cloud PubSub publish & consume

`/ea_monitor/` Latest EA Gauge Values
-------------------------------------

[![Docker Repository on Quay](https://quay.io/repository/rainchasers/ea_monitor/status "Docker Repository on Quay")](https://quay.io/repository/rainchasers/ea_monitor)

Daemon that polls the [EA Real Time Data API](http://environment.data.gov.uk/flood-monitoring/doc/reference) and publishes a stream of the latest reading values to a PubSub topic. The daemon executes in 3 phases:

1. *Station Discovery* (`discover.go`): paginate through the station summaries to gather the full list of stations.  
2. *Station Metadata* (`detail.go`): for each station get the required meta-data. The need for the Lg/Lat means an individual request needs to be made for each station.
3. *Measurement Update* (`update.go`): paginates through all measurements and publishes the latest measurement combined (`combine.go`) with the already gathered meta-data.

A standard API http request library (`http.go`) throttles the requests into the EA API.  