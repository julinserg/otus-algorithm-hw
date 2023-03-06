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

type SorterQuick struct{}

func (s *SorterQuick) split(array []int, L, R int) int {
	P := array[R]
	m := L - 1
	for i := L; i <= R; i++ {
		if array[i] <= P {
			m++
			swap(array, m, i)
		}
	}
	return m
}

func (s *SorterQuick) qsort(array []int, L, R int) {
	if L >= R {
		return
	}
	M := s.split(array, L, R)
	s.qsort(array, L, M-1)
	s.qsort(array, M+1, R)
}

func (s *SorterQuick) Sort(in []int) []int {
	N := len(in)
	s.qsort(in, 0, N-1)
	return in
}

func (s *SorterQuick) Name() string {
	return "Quick"
}

type SorterMerge struct{}

func (s *SorterMerge) merge(array []int, L, M, R int) {
	T := make([]int, R-L+1)
	a := L
	b := M + 1
	t := 0

	for a <= M && b <= R {
		if array[a] <= array[b] {
			T[t] = array[a]
			a++
		} else {
			T[t] = array[b]
			b++
		}
		t++
	}
	for a <= M {
		T[t] = array[a]
		t++
		a++
	}
	for b <= R {
		T[t] = array[b]
		t++
		b++
	}
	for i := L; i <= R; i++ {
		array[i] = T[i-L]
	}
}
func (s *SorterMerge) msort(array []int, L, R int) {
	if L >= R {
		return
	}
	M := (L + R) / 2
	s.msort(array, L, M)
	s.msort(array, M+1, R)
	s.merge(array, L, M, R)
}

func (s *SorterMerge) Sort(in []int) []int {
	N := len(in)
	s.msort(in, 0, N-1)
	return in
}

func (s *SorterMerge) Name() string {
	return "Merge"
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

	listSorterAlgo := []ISorter{&SorterQuick{}}
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
