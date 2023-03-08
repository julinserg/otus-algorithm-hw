package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
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
	fileList       []*os.File
	limitArraySize int
	fileA          *os.File
	fileB          *os.File
	fileC          *os.File
	fileD          *os.File
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

func (s *SorterExternal) GenerateTestData(numFile int, numLine int, maxNumber int) {
	s.limitArraySize = numLine
	s.generateTestFiles(numFile, numLine, maxNumber)
}

func (s *SorterExternal) RemoveTestData() {
	for _, el := range s.fileList {
		os.Remove(el.Name())
	}
}

func (s *SorterExternal) RemoveTempFile() {
	os.Remove(s.fileA.Name())
	os.Remove(s.fileB.Name())
	os.Remove(s.fileC.Name())
	os.Remove(s.fileD.Name())
}

func (s *SorterExternal) CreateTempFile() {
	var err error
	s.fileA, err = ioutil.TempFile("", "otus")
	if err != nil {
		log.Fatal(err)
	}
	s.fileB, err = ioutil.TempFile("", "otus")
	if err != nil {
		log.Fatal(err)
	}
	s.fileC, err = ioutil.TempFile("", "otus")
	if err != nil {
		log.Fatal(err)
	}
	s.fileD, err = ioutil.TempFile("", "otus")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SorterExternal) fileToArray(file *os.File) []int {
	result := make([]int, 0, s.limitArraySize)
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	for i := 0; i < s.limitArraySize; i++ {
		l, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		intVar, err := strconv.Atoi(string(l))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, intVar)
	}
	return result
}

func (s *SorterExternal) arrayToFile(array []int, file *os.File, isBegin bool) {
	if isBegin {
		os.Truncate(file.Name(), 0)
		file.Seek(0, 0)
	}

	for _, el := range array {
		file.WriteString(strconv.Itoa(el) + "\n")
	}
}

func (s *SorterExternal) splitSrcDataByTwoFileAB() {
	for index, el := range s.fileList {
		ar := s.fileToArray(el)
		sort.Ints(ar)
		isBegin := false
		if index == 0 {
			isBegin = true
		}
		if index%2 == 0 {
			s.arrayToFile(ar, s.fileA, isBegin)
		} else {
			s.arrayToFile(ar, s.fileB, isBegin)
		}
	}
}

func (s *SorterExternal) oneMerge(srcFile1, srcFile2, dstFile1, dstFile2 *os.File) {
	readerOne := bufio.NewReader(srcFile1)
	readerTwo := bufio.NewReader(srcFile2)
	writerOne := bufio.NewWriter(dstFile1)
	writerTwo := bufio.NewWriter(dstFile2)
	srcFile1.Seek(0, 0)
	srcFile2.Seek(0, 0)
	dstFile1.Seek(0, 0)
	dstFile2.Seek(0, 0)
	var valA, valB, oldValA, oldValB int
	isNextA := true
	isNextB := true
	count := 0
	outputWriter := writerOne
	currentWriterIsOne := true
	isBegin := true
	isFileAEndFirst := true
	isChangeA := false
	isChangeB := false
	for {
		if isNextA {
			strA, _, err := readerOne.ReadLine()
			if err != nil {
				isFileAEndFirst = true
				break
			}
			valA, _ = strconv.Atoi(string(strA))
		}
		if isNextB {
			strB, _, err := readerTwo.ReadLine()
			if err != nil {
				isFileAEndFirst = false
				break
			}
			valB, _ = strconv.Atoi(string(strB))
		}

		if oldValA > valA {
			isChangeA = true
		}
		if oldValB > valB {
			isChangeB = true
		}

		if isChangeA && isChangeB && !isBegin {
			if currentWriterIsOne {
				outputWriter = writerTwo
			} else {
				outputWriter = writerOne
			}
			currentWriterIsOne = !currentWriterIsOne
			isChangeA = false
			isChangeB = false
		}

		if !isChangeA && !isChangeB {
			if valA <= valB {
				outputWriter.WriteString(strconv.Itoa(valA) + "\n")
				isNextA = true
				isNextB = false
				fmt.Println("!!!!!!!!!!!!", valA, valB, valA, currentWriterIsOne)
			} else {
				outputWriter.WriteString(strconv.Itoa(valB) + "\n")
				isNextA = false
				isNextB = true
				fmt.Println("!!!!!!!!!!!!", valA, valB, valB, currentWriterIsOne)
			}
		} else if isChangeA && !isChangeB {
			outputWriter.WriteString(strconv.Itoa(valB) + "\n")
			isNextA = false
			isNextB = true
			fmt.Println("!!!!!!!!!!!!", valA, valB, valB, currentWriterIsOne)
		} else if isChangeB && !isChangeA {
			outputWriter.WriteString(strconv.Itoa(valA) + "\n")
			isNextA = true
			isNextB = false
			fmt.Println("!!!!!!!!!!!!", valA, valB, valA, currentWriterIsOne)
		} else if isChangeA && isChangeB {
			fmt.Println("------------", valA, valB, valA, currentWriterIsOne)
		}

		oldValA = valA
		oldValB = valB
		isBegin = false
		count++
	}
	if isFileAEndFirst {
		fmt.Println("File A end")
	} else {
		fmt.Println("File B end")
	}

	readerLost := readerOne
	if isFileAEndFirst {
		readerLost = readerTwo
		outputWriter.WriteString(strconv.Itoa(valB) + "\n")
		fmt.Println("!!!!!!!!!!!!", valB, currentWriterIsOne)
	} else {
		outputWriter.WriteString(strconv.Itoa(valA) + "\n")
		fmt.Println("!!!!!!!!!!!!", valA, currentWriterIsOne)
	}

	count++

	for {
		str, _, err := readerLost.ReadLine()
		if err != nil {
			break
		}
		val, _ := strconv.Atoi(string(str))
		outputWriter.WriteString(strconv.Itoa(val) + "\n")
		fmt.Println("!!!!!!!!!!!!", val, currentWriterIsOne)
		count++
	}
	fmt.Println("!!!!!!!!!!!!", count)
	writerOne.Flush()
	writerTwo.Flush()
	os.Truncate(srcFile1.Name(), 0)
	os.Truncate(srcFile2.Name(), 0)
}

func (s *SorterExternal) Sort() []int {
	s.CreateTempFile()
	s.splitSrcDataByTwoFileAB()
	fileReadOne := s.fileA
	fileReadTwo := s.fileB
	fileWriteOne := s.fileC
	fileWriteTwo := s.fileD

	iterateCount := 0
	for {
		s.oneMerge(fileReadOne, fileReadTwo, fileWriteOne, fileWriteTwo)
		s.PrintDstFiles()
		iterateCount++
		tempReadOne := fileReadOne
		fileReadOne = fileWriteOne
		fileWriteOne = tempReadOne
		tempReadTwo := fileReadTwo
		fileReadTwo = fileWriteTwo
		fileWriteTwo = tempReadTwo
		f1, _ := os.Stat(fileReadOne.Name())
		f2, _ := os.Stat(fileReadTwo.Name())
		if f1.Size() == 0 || f2.Size() == 0 {
			break
		}
	}
	fmt.Println("Iterate count", iterateCount)

	return nil
}

func (s *SorterExternal) PrintSrcFiles() {
	countSrc := 0
	for _, el := range s.fileList {
		el.Seek(0, 0)
		scanner := bufio.NewScanner(el)
		for scanner.Scan() {
			countSrc++
			fmt.Printf(string(scanner.Text()) + " ")
		}
	}
	fmt.Printf("\n")
	fmt.Println("Count in Src File", countSrc)
}

func (s *SorterExternal) PrintDstFiles() {
	fmt.Println("File A")
	s.fileA.Seek(0, 0)
	scanner := bufio.NewScanner(s.fileA)
	countA := 0
	for scanner.Scan() {
		countA++
		fmt.Printf(string(scanner.Text()) + " ")
	}
	fmt.Printf("\n")
	fmt.Println("Count in File A", countA)

	fmt.Println("File B")
	s.fileB.Seek(0, 0)
	scanner = bufio.NewScanner(s.fileB)
	countB := 0
	for scanner.Scan() {
		countB++
		fmt.Printf(string(scanner.Text()) + " ")
	}
	fmt.Printf("\n")
	fmt.Println("Count in File B", countB)

	fmt.Println("File C")
	s.fileC.Seek(0, 0)
	scanner = bufio.NewScanner(s.fileC)
	countC := 0
	for scanner.Scan() {
		countC++
		fmt.Printf(string(scanner.Text()) + " ")
	}
	fmt.Printf("\n")
	fmt.Println("Count in File C", countC)
	fmt.Println("File D")
	s.fileD.Seek(0, 0)
	scanner = bufio.NewScanner(s.fileD)
	countD := 0
	for scanner.Scan() {
		countD++
		fmt.Printf(string(scanner.Text()) + " ")
	}
	fmt.Printf("\n")
	fmt.Println("Count in File D", countD)
}

//.......................................................

var dir string
var isExternal bool

func init() {
	flag.StringVar(&dir, "dir", "", "dir tests")
	flag.BoolVar(&isExternal, "e", false, "is external sort")
}

type TestData struct {
	input, output []int
	sizeArray     int
}

func main() {
	flag.Parse()
	if isExternal {
		sortExt := &SorterExternal{}
		sortExt.GenerateTestData(10, 10, 10)
		defer sortExt.RemoveTestData()
		defer sortExt.RemoveTempFile()
		fmt.Println("Print src array")
		sortExt.PrintSrcFiles()
		sortExt.Sort()
		fmt.Println("\n Print result array")
		sortExt.PrintDstFiles()
	} else {
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
		}
	}

}
