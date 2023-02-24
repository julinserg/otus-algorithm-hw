package main

import (
	"flag"
)

type SorterBubble struct{}

func (s *SorterBubble) Sort(in []int) []int {
	result := make([]int, 0)
	return result
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
	testData, err := readTestData(dir + "\\0.random")
	if err != nil {
		panic(err)
	}
	s := &SorterBubble{}
	err = runTests(testData, s)
	if err != nil {
		panic(err)
	}
}
