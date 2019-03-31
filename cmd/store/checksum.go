package main

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
)

// Checksum provides a checksum of current state
func checksum(v interface{}) string {
	b := sha1.New()
	gob.NewEncoder(b).Encode(v)
	return hex.EncodeToString(b.Sum(nil))
}
