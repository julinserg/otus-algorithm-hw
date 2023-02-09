package main

import (
	"fmt"
	"strconv"
)

type ISolver interface {
	Solve(data []string) int
}

type Solver struct{}

func (s *Solver) solveLocal(numIterate int, sumA int, sumB int, count int) int {
	if numIterate == 0 {
		if sumA == sumB {
			count++
		}
		return count
	}
	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			count = s.solveLocal(numIterate-1, sumA+a, sumB+b, count)
		}
	}
	return count
}

func (s *Solver) Solve(data []string) int {
	numberDigit, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	count := s.solveLocal(numberDigit/2, 0, 0, 0)
	return count
}

func main() {
	var s Solver
	dataInput := make([]string, 0)
	dataInput = append(dataInput, "6")
	fmt.Printf("result %d\n", s.Solve(dataInput))
}
