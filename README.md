# Rainchasers Gauge Service

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub and presenting out a gauge API.

## [EA Flooding Monitoring API](http://environment.data.gov.uk/flood-monitoring/doc/reference)

- Recent levels polling in `/cmd/ea`, daily batch reconciliation via `/cmd/eaday`
- Station identifiers include an `@id` of the data URL, `RLOIid` and `wiskiID` also available

The [EA Hydrology API](https://environment.data.gov.uk/hydrology/doc/reference) provides access to quality checked historical data. This is not used yet.

## [NRW Levels API](https://api-portal.naturalresources.wales/docs/services/577521aed81b570928363e10/operations/577521e0d81b570928363e11)

- Recent levels polled in `/cmd/nrw`
- Station identifier is `RLOIid`

This is an authenticated API and requires an API key to be stored in k8s:

    kubectl create secret generic -n default nrw-apikey \
    --from-literal=key=<your API Key>

## [SEPA CSV Data](http://apps.sepa.org.uk/waterlevels/)

- Recent levels polled in `/cmd/sepa`
- Station identifiers is an integer "location code" which appears in the data URL

## Not Used Yet

- [SEPA Rainfall Data](https://www2.sepa.org.uk/rainfall/)
- [Other SEPA Datasets](https://www.sepa.org.uk/environment/environmental-data/)
- [EA Catchment Data API](https://environment.data.gov.uk/catchment-planning/ui/reference)

## Deployment

Deployed onto k8s (GKE), with a continuous deliovery pipeline via Google Cloud Build.

- Create a Google Pub/Sub topic `gauge` for the snapshot updates
- Create a service account that can access Pub/Sub topic for the applications
- Add the relevant GKE secrets for Honeycomb, NRW & Algolia
- Create Google Cloud build deployments for each application using the templated `cloudbuild.yaml` and the `_APP` subsitution variable.
- Deploy to k8s cluster using the relavnt `deployment.yaml`

### Pub/Sub & Firestore Permissions via Service Account

Following [this process to allow access to pub/sub from the deamons](https://cloud.google.com/kubernetes-engine/docs/tutorials/authenticating-to-cloud-platform.

- Create a rainchasers-app [service account](https://console.cloud.google.com/iam-admin/serviceaccounts) in the project. Grant the `Pub/Sub Editor` and `Cloud Datastore User` roles.
- load the json key file as a secret `kubectl create secret generic service-accn-key --from-file=key.json=<filename>.json`

### Honeycomb

The k8s cluster needs the Honeycomb API key in a secrets store:

    kubectl create secret generic -n default honeycomb-writekey \
    --from-literal=key=<your API Key>

Daemons will post events to Honeycomb alongside JSON structured logs to Stdout.

### Algolia

The `firestore` daemon persists the river state infomation to Algolia SAAS search service
and this requires algolia credentials:

kubectl create secret generic -n default algolia-writekey \
 --from-literal=id=<app ID> --from-literal=key=<admin API key>

### GKE

Create the k8s deployments:

    # setup kubectl access
    gcloud config set project rainchasers
    gcloud config set compute/zone europe-west2-b
    gcloud container clusters get-credentials prod

    # create the deployments
    kubectl create -f ./cmd/ea/deployment.yaml
    kubectl create -f ./cmd/nrw/deployment.yaml
    kubectl create -f ./cmd/sepa/deployment.yaml
    kubectl create -f ./cmd/store/deployment.yaml
