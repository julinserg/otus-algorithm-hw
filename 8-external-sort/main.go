package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
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

//................External sort..........................

type SorterExternal struct {
	fileList []*os.File
}

func (s *SorterExternal) testFileGenerate(numLine int, maxNumber int) *os.File {
	file, err := ioutil.TempFile("", "otus")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < numLine; i++ {
		file.WriteString(strconv.Itoa(rand.Intn(maxNumber)) + "\n")
	}
	return file
}

func (s *SorterExternal) generateTestFiles(numFile int, numLine int, maxNumber int) {
	s.fileList = make([]*os.File, 0, numFile)
	for i := 0; i < numFile; i++ {
		f := s.testFileGenerate(numLine, maxNumber)
		s.fileList = append(s.fileList, f)
	}
}

func (s *SorterExternal) Name() string {
	return "External"
}

func (s *SorterExternal) GenerateData(numFile int, numLine int, maxNumber int) {
	s.generateTestFiles(numFile, numLine, maxNumber)
}

func (s *SorterExternal) Sort() []int {

	return nil
}

//.......................................................

var dir string

func init() {
	flag.StringVar(&dir, "dir", "", "dir tests")
}

type TestData struct {
	input, output []int
	sizeArray     int
}

func main() {
	sortExt := &SorterExternal{}
	sortExt.GenerateData(10, 10, 10)
	for _, el := range sortExt.fileList {
		el.Seek(0, 0)
		r := bufio.NewReader(el)
		for i := 0; i < 10; i++ {
			l, _, err := r.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf(string(l) + " ")
		}
	}
	/*flag.Parse()
	listFolder := []string{"0.random", "1.digits", "2.sorted", "3.revers"}

	listSorterAlgo := []ISorter{&SorterMerge{}}
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
	}*/
}
