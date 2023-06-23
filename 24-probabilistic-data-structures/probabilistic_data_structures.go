package p24probabilisticdatastructures

import (
	"hash/crc32"

	"github.com/dustin/go-probably"
)

func countUniqueHyperLogLog(data []string) uint64 {
	hll := probably.NewHyperLogLog(0.0001)
	for _, d := range data {
		hll.Add(crc32.ChecksumIEEE([]byte(d)))
	}
	return hll.Count()
}

func countUniqueCountMinSketch(data []string, query string) uint64 {
	sk := probably.NewSketch(1<<20, 5)
	for _, d := range data {
		sk.Increment(d)
	}
	return uint64(sk.Count(query))
}
