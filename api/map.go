package main

import (
	"encoding/json"
	"io"
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
				text: 'R', 
				title: item.name, 
				subTitle: item.url });
			map.entities.push(pp);
		}
	}
	</script>
	<script type='text/javascript' src='https://www.bing.com/api/maps/mapcontrol?callback=loadMap' async defer></script>
</body>
</html>`

func writeCatalogueHTML(h *Handler, w io.Writer) error {
	w.Write([]byte(mapPreHTML))

	type point struct {
		UUID string  `json:"uuid"`
		Name string  `json:"name"`
		URL  string  `json:"url"`
		Lat  float32 `json:"lat"`
		Lng  float32 `json:"lng"`
	}

	pts := make([]point, 0, h.Rivers.Count())
	h.Rivers.Each(func(s Section) bool {
		pts = append(pts, point{
			UUID: s.UUID,
			Name: s.SectionName + ", " + s.RiverName,
			URL:  "todo",
			Lat:  s.Putin.Lat,
			Lng:  s.Putin.Lng,
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
