package p24probabilisticdatastructures

import (
	"testing"
)

type CountMinSketchTestData struct {
	id      int
	inSet   []string
	inQuery string
	out     uint64
}

var countMinSketchTestData = []CountMinSketchTestData{
	{1, []string{
		"1",
		"1",
		"1",
		"1",
		"1",
	}, "1", 5},
	{1, []string{
		"1",
		"1",
		"1",
		"1",
		"2",
	}, "1", 4},
	{1, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}, "1", 1},
	{1, []string{
		"1",
		"2",
		"1",
		"2",
		"1",
	}, "1", 3},
}

func runCountMinSketchTests(t *testing.T, f func(data []string, query string) uint64, funcName string, testCases []CountMinSketchTestData) {
	for _, test := range testCases {
		actual := f(test.inSet, test.inQuery)
		if actual != test.out {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.out)
		}
	}
}

func TestCountMinSketch(t *testing.T) {
	runCountMinSketchTests(t, countUniqueCountMinSketch, "countUniqueCountMinSketch", countMinSketchTestData)
}
