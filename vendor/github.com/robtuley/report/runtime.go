package report

import (
	"runtime"
	"time"
)

// RuntimeStatsEvery logs an event with golang runtime statistics at the specified time interval
func RuntimeStatsEvery(every time.Duration) {
	go doRuntimeStats(every)
}

func doRuntimeStats(every time.Duration) {
	ticker := time.NewTicker(every)
	memStats := &runtime.MemStats{}
	lastSampleTime := time.Now()
	var lastPauseNs uint64
	var lastNumGc uint32
	nsInMs := float64(time.Millisecond)

	for {
		select {
		case <-channel.IsDraining:
			return
		case <-ticker.C:
		}

		runtime.ReadMemStats(memStats)
		now := time.Now()

		Info("runtime.memory", Data{
			"allocated":  float64(memStats.Alloc),
			"mallocs":    float64(memStats.Mallocs),
			"frees":      float64(memStats.Frees),
			"heap":       float64(memStats.HeapAlloc),
			"stack":      float64(memStats.StackInuse),
			"goroutines": float64(runtime.NumGoroutine()),
		})

		gcData := Data{
			"pause_ms_total": float64(memStats.PauseTotalNs) / nsInMs,
		}

		if lastPauseNs > 0 {
			pauseSinceLastSample := memStats.PauseTotalNs - lastPauseNs
			gcData["pause_ms_per_second"] = float64(pauseSinceLastSample) / nsInMs / every.Seconds()
		}
		lastPauseNs = memStats.PauseTotalNs

		countGc := int(memStats.NumGC - lastNumGc)
		if lastNumGc > 0 {
			diff := float64(countGc)
			diffTime := now.Sub(lastSampleTime).Seconds()
			gcData["gc_per_second"] = diff / diffTime
		}

		if countGc > 0 {
			if countGc > 256 {
				Action("runtime.gc.error", Data{
					"error":   "missing some gc pause times",
					"countGc": countGc,
				})
				countGc = 256
			}

			var maxPause float64
			for i := 0; i < countGc; i++ {
				idx := int((memStats.NumGC-uint32(i))+255) % 256
				pause := float64(memStats.PauseNs[idx]) / nsInMs
				if pause > maxPause {
					maxPause = pause
				}
			}
			gcData["pause_max_ms"] = maxPause
		}

		Info("runtime.gc", gcData)

		lastNumGc = memStats.NumGC
		lastSampleTime = now
	}
}
