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
	return "Bubble"
}

type SorterInsert struct{}

func (s *SorterInsert) Sort(in []int) []int {
	N := len(in)
	for i := 1; i < N; i++ {
		for j := i; j > 0; j-- {
			if in[j] < in[j-1] {
				swap(&in, j, j-1)
			} else {
				break
			}
		}
	}
	return in
}

func (s *SorterInsert) Name() string {
	return "Insert"
}

type SorterShell struct{}

func (s *SorterShell) Sort(in []int) []int {
	N := len(in)
	for shift := N / 2; shift >= 1; shift = shift / 2 {
		for i := shift; i < N; i++ {
			for j := i; j >= shift; j -= shift {
				if in[j] < in[j-shift] {
					swap(&in, j, j-shift)
				} else {
					break
				}
			}
		}
	}

	return in
}

func (s *SorterShell) Name() string {
	return "Shell"
}

type SorterInsertShift struct{}

func (s *SorterInsertShift) Sort(in []int) []int {
	N := len(in)
	for i := 1; i < N; i++ {
		el := in[i]
		var j int
		for j = i; j > 0; j-- {
			if el < in[j-1] {
				in[j] = in[j-1]
			} else {
				break
			}
		}
		in[j] = el
	}
	return in
}

func (s *SorterInsertShift) Name() string {
	return "InsertShift"
}

type SorterInsertBinarySearch struct{}

func binarySearch(array *[]int, key int, left, right int) int {
	for left <= right {
		mid := (right + left) / 2
		if (*array)[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func (s *SorterInsertBinarySearch) Sort(in []int) []int {
	N := len(in)
	for i := 1; i < N; i++ {
		if in[i] < in[i-1] {
			el := in[i]
			k := binarySearch(&in, el, 0, i-1)
			var j int
			for j = i; j > k; j-- {
				in[j] = in[j-1]
			}
			in[j] = el
		}
	}
	return in
}

func (s *SorterInsertBinarySearch) Name() string {
	return "InsertBinarySearch"
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
	listSorterAlgo := []ISorter{&SorterBubble{}, &SorterInsert{}, &SorterShell{}, &SorterInsertShift{}, &SorterInsertBinarySearch{}}
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
