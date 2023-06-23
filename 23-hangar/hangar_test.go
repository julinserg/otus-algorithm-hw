package p23hangar

import (
	"testing"
)

type HangarTestData struct {
	id  int
	in  [][]int
	out int
}

var hangarTestData = []HangarTestData{
	{1, [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}, 12},
	{2, [][]int{
		{0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{1, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}, 8},
	{3, [][]int{
		{0, 0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
	}, 8},
	{4, [][]int{
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1, 0, 0},
	}, 9},
}

func runHangarTests(t *testing.T, f func(garden [][]int) int, funcName string, testCases []HangarTestData) {
	for _, test := range testCases {
		actual := f(test.in)
		if actual != test.out {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.out)
		}
	}
}

func TestHangarBruteForce(t *testing.T) {
	runHangarTests(t, hangarBruteForce, "hangarBruteForce", hangarTestData)
}

func TestHangarBruteForceWithOptimization(t *testing.T) {
	runHangarTests(t, hangarBruteForceWithOptimization, "hangarBruteForceWithOptimization", hangarTestData)
}

func TestHangarO3(t *testing.T) {
	runHangarTests(t, hangarO3, "hangarO3", hangarTestData)
}

func TestHangarO2(t *testing.T) {
	runHangarTests(t, hangarO2, "hangarO2", hangarTestData)
}
