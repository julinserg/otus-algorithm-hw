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

type SorterSelection struct{}

func (s *SorterSelection) Sort(in []int) []int {
	N := len(in)
	for i := N - 1; i > 0; i-- {
		maxIdx := 0
		for j := 0; j <= i; j++ {
			if in[j] > in[maxIdx] {
				maxIdx = j
			}
		}
		swap(&in, i, maxIdx)
	}
	return in
}

func (s *SorterSelection) Name() string {
	return "Selection"
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
	listSorterAlgo := []ISorter{&SorterSelection{}}
	for _, lf := range listFolder {
		log.Printf("Test folder - %s \n", lf)
		testData, err := readTestData(dir+"\\"+lf, 6)
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
