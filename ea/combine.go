package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func applyUpdatesToRefSnaps(
	refC <-chan gauge.Snapshot,
	inLatestC <-chan gauge.SnapshotUpdate,
	inHistoryC <-chan gauge.SnapshotUpdate,
) (<-chan gauge.Snapshot, <-chan gauge.Snapshot) {

	outLatestC := make(chan gauge.Snapshot)
	outHistoryC := make(chan gauge.Snapshot)

	go func() {
		ref := make(map[string]gauge.Snapshot, 10000)

		for {
			select {
			case s := <-refC:
				ref[s.Url] = s
				outLatestC <- s
			case u := <-inLatestC:
				tpl, exists := ref[u.Url]
				if exists {
					outLatestC <- tpl.Apply(u)
				}
			case u := <-inHistoryC:
				tpl, exists := ref[u.Url]
				if exists {
					outHistoryC <- tpl.Apply(u)
				}
			}
		}
	}()

	return outLatestC, outHistoryC
}
