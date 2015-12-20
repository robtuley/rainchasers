package main

func applyUpdatesToRefSnaps(refSnapC chan Snapshot, updateSnapC chan SnapshotUpdate) chan Snapshot {
	pubSnapC := make(chan Snapshot)

	go func() {
		ref := make(map[string]Snapshot, 5000)

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
