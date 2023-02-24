package main

import (
	"flag"
	"fmt"
)

func swap(array *[]int, indexA, indexB int) {
	t := (*array)[indexA]
	(*array)[indexA] = (*array)[indexB]
	(*array)[indexB] = t
}

type SorterBubble struct{}

func (s *SorterBubble) Sort(in []int) []int {
	N := len(in)
	for i := 0; i < N; i++ {
		for j := 0; j < N-1-i; j++ {
			if in[j] > in[j+1] {
				swap(&in, j, j+1)
			}
		}
	}
	return in
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
