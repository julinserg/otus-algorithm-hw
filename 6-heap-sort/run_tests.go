package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type ISorter interface {
	Sort(in []int) []int
}

func arrayStrToInt(array []string) []int {
	inputArrayInt := make([]int, 0)
	for _, i := range array {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputArrayInt = append(inputArrayInt, j)
	}
	return inputArrayInt
}

func readTestData(dirPath string) ([]TestData, error) {
	_, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	result := make([]TestData, 0)
	for numTest := 0; numTest < 7; numTest++ {

		fileIn, err := os.Open(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".in"))
		if err != nil {
			break
		}
		defer fileIn.Close()

		readerIn := bufio.NewReader(fileIn)
		line, _, err := readerIn.ReadLine()
		if err != nil {
			panic(err)
		}
		sizeArray, err := strconv.Atoi(string(line))
		if err != nil {
			panic(err)
		}
		line, _, err = readerIn.ReadLine()
		if err != nil {
			panic(err)
		}
		inputArrayStr := strings.Split(string(line), " ")
		inputArrayInt := arrayStrToInt(inputArrayStr)

		fileOut, err := os.Open(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".out"))
		if err != nil {
			break
		}
		defer fileOut.Close()

		readerOut := bufio.NewReader(fileOut)
		line, _, err = readerOut.ReadLine()
		if err != nil {
			panic(err)
		}
		outputArrayStr := strings.Split(string(line), " ")
		outputArrayInt := arrayStrToInt(outputArrayStr)

		result = append(result, TestData{input: inputArrayInt, output: outputArrayInt, sizeArray: sizeArray})

	}
	return result, nil
}

func runTests(testData []TestData, s ISorter) error {
	for _, data := range testData {
		result := s.Sort(data.input)

		if reflect.DeepEqual(result, data.output) {
			fmt.Printf("%d - Success \n", data.sizeArray)
		} else {
			fmt.Printf("%d - Fail. Expected: %v. Actual: %v \n", data.sizeArray, data.output, result)
		}
	}
	return nil
}
