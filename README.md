Rainchasers Gauge Service
=========================

Responsible for retrieving updates from various flavours of river gauges and providing a consistent stream of these through Google Pub/Sub, archiving into Google BigQuery and presenting out a HTML gauge view service for the Rainchasers site.

`/ea_monitor/` Latest Gauge View
--------------------------------

to-do:

publish to pub/sub -- check encoder
controlled shutdown aftr x, max gauges, etc
docker wrapper

to raise with beta program:

-- pagination of stations based on measures
-- measures key is both an array OR a struct in station detail. So annoying.
-- http://environment.data.gov.uk/flood-monitoring/id/measures/1415TH-level-stage-i-15_min-mASD 
has an array as a value
-- I need to poll all stations individually to get lat/long, worth considering bringing this into list summary. 
-- lat/lg arrays! http://environment.data.gov.uk/flood-monitoring/id/stations/L1322



