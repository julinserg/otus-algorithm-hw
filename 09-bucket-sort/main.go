package main

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"flag"
	"log"
)

func findMax(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

type SorterCounting struct{}

func (s *SorterCounting) Sort(arr []int) []int {
	max := findMax(arr)
	counter := make([]int, max+1)

	for _, e := range arr {
		counter[e] += 1
	}

	res := make([]int, len(arr))

	b := 0
	for i := 0; i < len(counter); i++ {
		for j := 0; j < counter[i]; j++ {
			res[b] = i
			b++
		}
	}

	return res
}

func (s *SorterCounting) Name() string {
	return "Counting"
}

type SorterRadix struct{}

const digit = 4
const maxbit = -1 << 31

func (s *SorterRadix) Sort(data []int) []int {
	// Implementation copy from https://www.golangprograms.com/golang-program-for-implementation-of-radix-sort.html
	// with small my modify and bug fix (covert from int to int32 and from int32 to int)
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(data))
	var g int32
	for i, e := range data {
		g = int32(e) ^ maxbit
		binary.Write(buf, binary.LittleEndian, g)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	result := make([]int, len(data))
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		result[i] = int(w) ^ maxbit
	}
	return result
}

func (s *SorterRadix) Name() string {
	return "Radix"
}

type SorterBucket struct{}

func (s *SorterBucket) Sort(arr []int) []int {
	N := len(arr)
	max := findMax(arr)

	buckets := make([]*list.List, N)
	for i := 0; i < N; i++ {
		b := arr[i] * N / (max + 1)
		if buckets[b] == nil {
			buckets[b] = list.New()
		}
		var itForInsert *list.Element
		for it := buckets[b].Front(); it != nil; it = it.Next() {
			isPotentialInsert := false
			if it.Next() == nil {
				isPotentialInsert = true
			} else if arr[i] < it.Next().Value.(int) {
				isPotentialInsert = true
			}
			if arr[i] >= it.Value.(int) && isPotentialInsert {
				itForInsert = it
			}
		}
		if itForInsert != nil {
			buckets[b].InsertAfter(arr[i], itForInsert)
		} else {
			buckets[b].PushFront(arr[i])
		}

	}
	result := make([]int, N)
	index := 0
	for i := 0; i < N; i++ {
		if buckets[i] == nil {
			continue
		}
		for it := buckets[i].Front(); it != nil; it = it.Next() {
			result[index] = it.Value.(int)
			index++
		}
	}
	return result
}

func (s *SorterBucket) Name() string {
	return "Bucket"
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
	listSorterAlgo := []ISorter{&SorterBucket{}}
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
