package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ISolver interface {
	Solve(numberDigits int64) int64
}

func readTestData(dirPath string) ([]TestData, error) {
	_, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	result := make([]TestData, 0)
	numTest := 0
	for {
		inputDataRaw, err := os.ReadFile(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".in"))
		if err != nil {
			break
		}
		outputDataRaw, err := os.ReadFile(filepath.Join(dirPath, "test."+strconv.Itoa(numTest)+".out"))
		if err != nil {
			break
		}
		input := string(inputDataRaw)
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
		output := string(outputDataRaw)
		output = strings.TrimSuffix(output, "\n")
		output = strings.TrimSuffix(output, "\r")
		result = append(result, TestData{input, output})
		numTest++
	}
	return result, nil
}

func runTests(testData []TestData, s ISolver) error {
	for _, data := range testData {

		inputInt, err := strconv.ParseInt(data.input, 10, 64)
		if err != nil {
			return err
		}
		outputInt, err := strconv.ParseInt(data.output, 10, 64)
		if err != nil {
			return err
		}
		res := s.Solve(2 * inputInt)

		if res == outputInt {
			fmt.Printf("%d - Success \n", res)
		} else {
			fmt.Printf("%d - Fail. Expected: %d. Actual: %d \n", res, outputInt, res)
		}
	}
	return nil
}
