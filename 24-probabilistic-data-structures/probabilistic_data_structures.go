package p24probabilisticdatastructures

import (
	"github.com/dustin/go-probably"
)

func countUniqueCountMinSketch(data []string, query string) uint64 {
	sk := probably.NewSketch(1<<20, 5)
	for _, d := range data {
		sk.Increment(d)
	}
	return uint64(sk.Count(query))
}
