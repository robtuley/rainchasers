package report

import (
	"runtime"
)

func runtimeData() Data {
	data := Data{}

	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	data["stack.mb"] = float64(m.StackSys) / float64(1024*1024)
	data["heap.mb"] = float64(m.HeapAlloc) / float64(1024*1024)
	data["goroutine.count"] = runtime.NumGoroutine()
	data["gc.pause.ns"] = m.PauseNs[(m.NumGC+255)%256]

	return data
}
