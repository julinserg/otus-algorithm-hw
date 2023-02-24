package main

import (
	"flag"
	"fmt"
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
	listFolder := []string{"0.random", "1.digits", "2.sorted", "3.revers"}
	for _, lf := range listFolder {
		fmt.Printf("Test folder - %s \n", lf)
		testData, err := readTestData(dir + "\\" + lf)
		if err != nil {
			panic(err)
		}
		s := &SorterBubble{}
		err = runTests(testData, s)
		if err != nil {
			panic(err)
		}
	}

}
