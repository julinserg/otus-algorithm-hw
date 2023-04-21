package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ISorter interface {
	Sort(in []int) []int
	Name() string
}

func arrayStrToInt(array []string) []int {
	inputArrayInt := make([]int, 0)
	for index, i := range array {
		if i == "" {
			continue
		}
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		_ = index
		inputArrayInt = append(inputArrayInt, j)
	}
	return inputArrayInt
}

func readTestData(dirPath string, levelLimit int) ([]TestData, error) {
	_, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	result := make([]TestData, 0)
	for numTest := 0; numTest < levelLimit; numTest++ {

		fileIn, err := os.Open(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".in"))
		if err != nil {
			panic(err)
		}
		defer fileIn.Close()
		//fmt.Printf("Open file in - %s \n", fileIn.Name())
		scannerIn := bufio.NewScanner(fileIn)
		bufIn := make([]byte, 0, 1024*1024*1024)
		scannerIn.Buffer(bufIn, 1024*1024*1024)
		scannerIn.Scan()
		if err := scannerIn.Err(); err != nil {
			log.Fatal(err)
		}
		sizeArray, err := strconv.Atoi(scannerIn.Text())
		if err != nil {
			panic(err)
		}
		inputArrayIntAll := make([]int, 0)
		for scannerIn.Scan() {
			inputArrayStr := strings.Split(scannerIn.Text(), " ")
			inputArrayInt := arrayStrToInt(inputArrayStr)
			inputArrayIntAll = append(inputArrayIntAll, inputArrayInt...)
		}
		if err := scannerIn.Err(); err != nil {
			panic(err)
		}

		fileOut, err := os.Open(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".out"))
		if err != nil {
			panic(err)
		}
		defer fileOut.Close()
		//fmt.Printf("Open file out - %s \n", fileOut.Name())
		outputArrayIntAll := make([]int, 0)
		scannerOut := bufio.NewScanner(fileOut)
		bufOut := make([]byte, 0, 1024*1024*1024)
		scannerOut.Buffer(bufOut, 1024*1024*1024)
		for scannerOut.Scan() {
			outputArrayStr := strings.Split(scannerOut.Text(), " ")
			outputArrayInt := arrayStrToInt(outputArrayStr)
			outputArrayIntAll = append(outputArrayIntAll, outputArrayInt...)
		}
		if err := scannerOut.Err(); err != nil {
			panic(err)
		}

		result = append(result, TestData{input: inputArrayIntAll, output: outputArrayIntAll, sizeArray: sizeArray})

	}
	return result, nil
}

func runTests(testData []TestData, s ISorter) error {
	for _, data := range testData {
		dataInput := make([]int, len(data.input))

		copy(dataInput, data.input)
		start := time.Now()
		result := s.Sort(dataInput)
		elapsed := time.Since(start)
		if reflect.DeepEqual(result, data.output) {
			log.Printf("(%s) %d - Success (time: %s) \n", s.Name(), data.sizeArray, elapsed)
		} else {
			if data.sizeArray <= 100 {
				log.Printf("(%s) %d - Fail. Expected: %v. Actual: %v \n", s.Name(), data.sizeArray, data.output, result)
			} else {
				log.Printf("(%s) %d - Fail. \n", s.Name(), data.sizeArray)
			}
		}
	}
	return nil
}
