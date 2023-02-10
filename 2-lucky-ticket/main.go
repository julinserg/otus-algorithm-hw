package main

import (
	"flag"
)

type SolverLuckyTicket struct {
	digit [4]int
}

func (s *SolverLuckyTicket) solveLocal(numIterate int64, sumA int64, sumB int64, count int64) int64 {
	if numIterate == 0 {
		isLucky := false
		if sumA == sumB {
			count++
			isLucky = true
		}
		if isLucky {
			//fmt.Printf("digit = %d %d %d %d - Lucky. \n", s.digit[0], s.digit[1], s.digit[2], s.digit[3])
		} else {
			//fmt.Printf("digit = %d %d %d %d \n", s.digit[0], s.digit[1], s.digit[2], s.digit[3])
		}

		return count
	}
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			if numIterate == 2 {
				s.digit[0] = a
				s.digit[2] = b
			} else {
				s.digit[1] = a
				s.digit[3] = b
			}

			count = s.solveLocal(numIterate-1, sumA+int64(a), sumB+int64(b), count)
		}
	}
	return count
}

func (s *SolverLuckyTicket) Solve(numberDigits int64) int64 {
	s.digit[0] = 0
	s.digit[1] = 0
	s.digit[2] = 0
	s.digit[3] = 0
	count := s.solveLocal(4/2, 0, 0, 0)
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
