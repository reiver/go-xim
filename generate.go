package xim

import (
	"math/rand"
	"time"
)

var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
)

func generate() uint64 {

	var now int64 = time.Now().Unix()
	return generateforunixtime(now)
}

func generateforunixtime(unixtimestamp int64) uint64 {
	var chaos uint64 = randomness.Uint64()

	var value uint64 = compile(uint64(unixtimestamp), chaos)

	return value

}
