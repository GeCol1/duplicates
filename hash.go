package main

import (
	"crypto/sha1"
	"fmt"
)

func GetHash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}
