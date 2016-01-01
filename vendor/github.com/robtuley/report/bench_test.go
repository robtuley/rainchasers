package report

import (
	"errors"
	"testing"
)

func BenchmarkInfo(b *testing.B) {
	Global(Data{"application": "myAppName"})
	for i := 0; i < b.N; i++ {
		Info("event.name", Data{"a": "aString", "z": 12})
	}
}

func BenchmarkAction(b *testing.B) {
	Global(Data{"application": "myAppName"})
	for i := 0; i < b.N; i++ {
		Action("event.name", Data{"a": "aString", "error": errors.New("a test").Error()})
	}
}
