[![Go Report Card](https://goreportcard.com/badge/github.com/rainchasers/com.rainchasers.gauge)](https://goreportcard.com/report/github.com/rainchasers/com.rainchasers.gauge)

# Rainchasers Gauge Service

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub and presenting out a gauge API.

## [EA Flooding Monitoring API](http://environment.data.gov.uk/flood-monitoring/doc/reference)

* Recent levels polling in `/cmd/ea`, daily batch reconciliation via `/cmd/eaday`
* Station identifiers include an `@id` of the data URL, `RLOIid` and `wiskiID` also available

The [EA Hydrology API](https://environment.data.gov.uk/hydrology/doc/reference) provides access to quality checked historical data. This is not used yet.

## [NRW Levels API](https://api-portal.naturalresources.wales/docs/services/577521aed81b570928363e10/operations/577521e0d81b570928363e11)

This is not yet used.

## [SEPA CSV Data](http://apps.sepa.org.uk/waterlevels/)

* Recent levels polled in `/cmd/sepa`
* Station identifiers is an integer "location code" which appears in the data URL

## [EA Catchment Data API](https://environment.data.gov.uk/catchment-planning/ui/reference)

This is not used yet.

## Honeycomb

The k8s cluster needs the Honeycomb API key in a secrets store:

    kubectl create secret generic -n default honeycomb-writekey \
    --from-literal=key=<your API Key>

Daemons will post events to Honeycomb alongside JSON structured logs to Stdout.
