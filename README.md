[![Go Report Card](https://goreportcard.com/badge/github.com/rainchasers/com.rainchasers.gauge)](https://goreportcard.com/report/github.com/rainchasers/com.rainchasers.gauge)

# Rainchasers Gauge Service

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub and presenting out a gauge API.

## Honeycomb

The k8s cluster needs the Honeycomb API key in a secrets store:

    kubectl create secret generic -n default honeycomb-writekey \
    --from-literal=key=<your API Key>

Daemons will post events to Honeycomb alongside JSON structured logs to Stdout.
