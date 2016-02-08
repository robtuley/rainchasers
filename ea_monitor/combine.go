package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func applyUpdatesToRefSnaps(refSnapC chan gauge.Snapshot,
	updateSnapC chan gauge.SnapshotUpdate) chan gauge.Snapshot {

	pubSnapC := make(chan gauge.Snapshot)

	go func() {
		ref := make(map[string]gauge.Snapshot, 10000)

		for {
			select {
			case r := <-refSnapC:
				ref[r.Url] = r
				pubSnapC <- r
			case u := <-updateSnapC:
				tpl, exists := ref[u.Url]
				if exists {
					pubSnapC <- tpl.Apply(u)
				}
			}
		}
	}()

	return pubSnapC
}
