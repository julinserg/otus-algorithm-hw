package main

import (
	"flag"
	"log"
)

func swap(array []int, indexA, indexB int) {
	t := array[indexA]
	array[indexA] = array[indexB]
	array[indexB] = t
}

type SorterHeap struct{}

func heapify(array []int, root, size int) {
	x := root
	l := 2*x + 1
	r := 2*x + 2
	if l < size && array[l] > array[x] {
		x = l
	}
	if r < size && array[r] > array[x] {
		x = r
	}

	if root == x {
		return
	} else {
		swap(array, x, root)
		heapify(array, x, size)
	}
}

func (s *SorterHeap) Sort(in []int) []int {
	N := len(in)
	for i := N/2 - 1; i >= 0; i-- {
		heapify(in, i, N)
	}
	for i := N - 1; i >= 0; i-- {
		swap(in, 0, i)
		heapify(in, 0, i)
	}
	return in
}

func (s *SorterHeap) Name() string {
	return "Heap"
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
		swap(in, i, maxIdx)
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
	listSorterAlgo := []ISorter{&SorterHeap{}}
	for _, lf := range listFolder {
		log.Printf("Test folder - %s \n", lf)
		testData, err := readTestData(dir+"\\"+lf, 7)
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
