package main

import (
	"strconv"
	"testing"
)

func TestDeDupeOverflow(t *testing.T) {
	dedup := newDeDupeCache(100)

	for i := 0; i < 200; i++ {
		k := strconv.Itoa(i)
		dedup.Set(k)
	}

	for i := 100; i < 200; i++ {
		k := strconv.Itoa(i)
		if !dedup.Exists(k) {
			t.Error("Key premature ejection", k)
		}
	}

	for i := 0; i < 99; i++ {
		k := strconv.Itoa(i)
		if dedup.Exists(k) {
			t.Error("Key kept too long", k)
		}
	}
}
