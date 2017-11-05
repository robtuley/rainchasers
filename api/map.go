package main

import (
	"encoding/json"
	"io"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

const mapPreHTML = `<!DOCTYPE html>
<html>
<head>
  <title>Explore Rainchasers Dataset</title>
  <meta http-equiv='Content-Type' content='text/html; charset=utf-8'/>
</head>
<body style="margin:0; padding:0">
	<script>
	var RainchasersCatalogue = `

const mapPostHTML = `
	</script>
	<div id=map style="width: 100vw; height: 100vh;"></div>
	<script>
	function loadMap() {
		var map = new Microsoft.Maps.Map(document.getElementById('map'),{
			credentials: 'AiPhuqF9yzxns1TkCbMG30W_lBn0x9yRBtnIEzFXZVfd_-AUKN9mPMGDab6Pra1M'
		});
		for (var i = 0, len = RainchasersCatalogue.length; i < len; i++) {
			var item = RainchasersCatalogue[i];
			var ll = new Microsoft.Maps.Location(item.lat, item.lng)
			var pp = new Microsoft.Maps.Pushpin(ll, { 
				color: (item.type=='R'?'red':'blue'),
				text: item.type, 
				title: item.name });
			pp.url = item.url
			Microsoft.Maps.Events.addHandler(pp, 'click', pushpinClicked);
			map.entities.push(pp);
		}

		function pushpinClicked(e) {
			if (e.target.url) {
				window.open(e.target.url, '_blank');
			}
		}
	}
	</script>
	<script type='text/javascript' src='https://www.bing.com/api/maps/mapcontrol?callback=loadMap' async defer></script>
</body>
</html>`

func writeCatalogueHTML(h *Handler, w io.Writer) error {
	w.Write([]byte(mapPreHTML))

	type point struct {
		Type string  `json:"type"`
		UUID string  `json:"uuid"`
		Name string  `json:"name"`
		URL  string  `json:"url"`
		Lat  float32 `json:"lat"`
		Lng  float32 `json:"lng"`
	}

	pts := make([]point, 0, h.Rivers.Count()+h.Gauge.Count())
	h.Rivers.Each(func(s Section) bool {
		pts = append(pts, point{
			Type: "R",
			UUID: s.UUID,
			Name: s.SectionName + ", " + s.RiverName,
			URL:  "todo",
			Lat:  s.Putin.Lat,
			Lng:  s.Putin.Lng,
		})
		return true
	})
	h.Gauge.Each(func(s *gauge.Snapshot) bool {
		pts = append(pts, point{
			Type: "G",
			UUID: s.Station.UUID(),
			Name: s.Station.Name,
			URL:  s.Station.DataURL,
			Lat:  s.Station.Lat,
			Lng:  s.Station.Lg,
		})
		return true
	})

	enc := json.NewEncoder(w)
	if err := enc.Encode(pts); err != nil {
		return err
	}

	w.Write([]byte(mapPostHTML))
	return nil
}
