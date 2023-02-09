package main

import (
	"flag"
)

type SolverLuckyTicket struct{}

func (s *SolverLuckyTicket) solveLocal(numIterate int64, sumA int64, sumB int64, count int64) int64 {
	if numIterate == 0 {
		if sumA == sumB {
			count++
		}
		return count
	}
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			count = s.solveLocal(numIterate-1, sumA+int64(a), sumB+int64(b), count)
		}
	}
	return count
}

func (s *SolverLuckyTicket) Solve(numberDigits int64) int64 {
	count := s.solveLocal(numberDigits/2, 0, 0, 0)
	return count
}

var dir string

func init() {
	flag.StringVar(&dir, "dir", "", "dir tests")
}

type TestData struct {
	input, output string
}

func main() {
	flag.Parse()
	testData, err := readTestData(dir)
	if err != nil {
		panic(err)
	}
	s := &SolverLuckyTicket{}
	err = runTests(testData, s)
	if err != nil {
		panic(err)
	}
}
