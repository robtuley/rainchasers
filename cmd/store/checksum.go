package main

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
)

// Checksum provides a checksum of current state
func checksum(values ...interface{}) string {
	b := sha1.New()
	encoder := gob.NewEncoder(b)
	for _, v := range values {
		encoder.Encode(v)
	}
	return hex.EncodeToString(b.Sum(nil))
}
