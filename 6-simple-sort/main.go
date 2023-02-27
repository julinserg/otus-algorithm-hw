package main

import (
	"flag"
	"log"
)

func swap(array *[]int, indexA, indexB int) {
	t := (*array)[indexA]
	(*array)[indexA] = (*array)[indexB]
	(*array)[indexB] = t
}

type SorterBubble struct{}

func (s *SorterBubble) Sort(in []int) []int {
	N := len(in)
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1-i; j++ {
			if in[j] > in[j+1] {
				swap(&in, j, j+1)
			}
		}
	}
	return in
}

func (s *SorterBubble) Name() string {
	return "SorterBubble"
}

type SorterInsert struct{}

func (s *SorterInsert) Sort(in []int) []int {
	N := len(in)
	for i := 1; i < N; i++ {
		for j := i; j > 0; j-- {
			if in[j] < in[j-1] {
				swap(&in, j, j-1)
			}
		}
	}
	return in
}

func (s *SorterInsert) Name() string {
	return "SorterInsert"
}

type SorterShell struct{}

func (s *SorterShell) Sort(in []int) []int {
	N := len(in)
	for i := 1; i < N; i++ {
		for j := i; j > 0; j-- {
			if in[j] < in[j-1] {
				swap(&in, j, j-1)
			}
		}
	}
	return in
}

func (s *SorterShell) Name() string {
	return "SorterShell"
}

var dir string

func init() {
	flag.StringVar(&dir, "dir", "", "dir tests")
}

type TestData struct {
	input, output []int
	sizeArray     int
}

func main() {
	flag.Parse()
	listFolder := []string{"0.random", "1.digits", "2.sorted", "3.revers"}
	listSorterAlgo := []ISorter{&SorterBubble{}, &SorterInsert{}, &SorterShell{}}
	for _, lf := range listFolder {
		log.Printf("Test folder - %s \n", lf)
		testData, err := readTestData(dir+"\\"+lf, 4)
		if err != nil {
			panic(err)
		}
		for _, alg := range listSorterAlgo {
			err = runTests(testData, alg)
			if err != nil {
				panic(err)
			}
		}
	}
}
