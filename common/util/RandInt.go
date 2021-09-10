package util

import (
	"math/rand"
	"time"
)

var s2 = rand.NewSource(time.Now().UnixNano())
var r2 = rand.New(s2)

// max is inclusive
func RandInt(max int) int {
	return r2.Intn(max + 1)
}