package iid

import (
	"math/rand"
	"time"
)

var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
)

func generate() uint64 {

	var now    int64 = time.Now().Unix()
	var chaos uint64 = randomness.Uint64()

	var value uint64 = compile(uint64(now), chaos)

	return value
}
