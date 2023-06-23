package p24probabilisticdatastructures

import (
	"testing"
)

type HyperLogLogTestData struct {
	id  int
	in  []string
	out uint64
}

var hyperLogLogTestData = []HyperLogLogTestData{
	{1, []string{
		"123",
		"123",
		"123",
		"123",
		"123",
	}, 0},
	{1, []string{
		"123",
		"123",
		"123",
		"123",
		"124",
	}, 2},
	{1, []string{
		"123",
		"124",
		"125",
		"126",
		"127",
	}, 5},
	{1, []string{
		"123",
		"123",
		"121",
		"122",
		"123",
	}, 3},
}

type CountMinSketchTestData struct {
	id      int
	inSet   []string
	inQuery string
	out     uint64
}

var countMinSketchTestData = []CountMinSketchTestData{
	{1, []string{
		"123",
		"123",
		"123",
		"123",
		"123",
	}, "123", 5},
	{1, []string{
		"123",
		"123",
		"123",
		"123",
		"124",
	}, "123", 4},
	{1, []string{
		"123",
		"124",
		"125",
		"126",
		"127",
	}, "123", 1},
	{1, []string{
		"123",
		"123",
		"121",
		"122",
		"123",
	}, "123", 3},
}

func runHyperLogLogTests(t *testing.T, f func(data []string) uint64, funcName string, testCases []HyperLogLogTestData) {
	for _, test := range testCases {
		actual := f(test.in)
		if actual != test.out {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.out)
		}
	}
}

func runCountMinSketchTests(t *testing.T, f func(data []string, query string) uint64, funcName string, testCases []CountMinSketchTestData) {
	for _, test := range testCases {
		actual := f(test.inSet, test.inQuery)
		if actual != test.out {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.out)
		}
	}
}

func TestHyperLogLog(t *testing.T) {
	runHyperLogLogTests(t, countUniqueHyperLogLog, "countUniqueHyperLogLog", hyperLogLogTestData)
}

func TestCountMinSketch(t *testing.T) {
	runCountMinSketchTests(t, countUniqueCountMinSketch, "countUniqueCountMinSketch", countMinSketchTestData)
}
