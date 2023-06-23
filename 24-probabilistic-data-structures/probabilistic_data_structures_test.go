package p24probabilisticdatastructures

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/dustin/go-probably"
	"github.com/stretchr/testify/require"
)

type CountMinSketchTestData struct {
	id      int
	inSet   []string
	inQuery string
	out     uint64
}

var countMinSketchTestData = []CountMinSketchTestData{
	{1, []string{
		"1",
		"1",
		"1",
		"1",
		"1",
	}, "1", 5},
	{1, []string{
		"1",
		"1",
		"1",
		"1",
		"2",
	}, "1", 4},
	{1, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}, "1", 1},
	{1, []string{
		"1",
		"2",
		"1",
		"2",
		"1",
	}, "1", 3},
}

func runCountMinSketchTests(t *testing.T, f func(data []string, query string) uint64, funcName string, testCases []CountMinSketchTestData) {
	for _, test := range testCases {
		actual := f(test.inSet, test.inQuery)
		if actual != test.out {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.out)
		}
	}
}

func countUniqueCountMinSketch(data []string, query string) uint64 {
	sk := probably.NewSketch(1<<20, 5)
	for _, d := range data {
		sk.Increment(d)
	}
	return uint64(sk.Count(query))
}

func TestCountMinSketch(t *testing.T) {
	runCountMinSketchTests(t, countUniqueCountMinSketch, "countUniqueCountMinSketch", countMinSketchTestData)
}

//*******************https://www.kaggle.com/datasets/datasnaek/youtube-new***********************************

type CountMinSketchTestYoutubeData struct {
	id       int
	inW      int
	inD      int
	outError int
}

var countMinSketchTestYoutubeData = []CountMinSketchTestYoutubeData{
	{1, 10000000, 5, 0},
	{2, 1000000, 5, 10},
	{3, 100000, 5, 49677},
	{4, 10000, 5, 157673},
	{5, 1000, 5, 157673},
	{6, 100, 5, 157673},
	{7, 10000000, 4, 0},
	{8, 10000000, 3, 0},
	{9, 10000000, 2, 30},
	{10, 10000000, 1, 2419},
	{10, 500000, 2, 11531},
}

func runCountMinSketchTestsYoutubeData(t *testing.T, f func(t *testing.T, w int, d int) int, funcName string, testCases []CountMinSketchTestYoutubeData) {
	for _, test := range testCases {
		actual := f(t, test.inW, test.inD)
		if actual != test.outError {
			t.Errorf("%s(id dataset=%d) = %v; want %v", funcName, test.id, actual, test.outError)
		}
	}
}

func getRealSizeOf(v interface{}) (int, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}

func countMinSketchOnYoutubeData(t *testing.T, w int, d int) int {
	// https://www.kaggle.com/datasets/datasnaek/youtube-new
	file, err := os.Open("RUvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	etalonMap := make(map[string]int)
	sk := probably.NewSketch(w, d)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParse := strings.Split(line, ",")
		tagsLine := lineParse[6]
		tagsLine = strings.ReplaceAll(tagsLine, "\"", "")
		tags := strings.Split(tagsLine, "|")
		for _, tag := range tags {
			sk.Increment(tag)
			etalonMap[tag] += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m, _ := getRealSizeOf(etalonMap)
	fmt.Println("Size map", m)

	require.Equal(t, 367, etalonMap["футбол"])
	require.Equal(t, 84, etalonMap["хоккей"])
	require.Equal(t, 43, etalonMap["бокс"])
	require.Equal(t, 33, etalonMap["смешанные единоборства"])
	require.Equal(t, 32, etalonMap["плавание"])
	require.Equal(t, 12, etalonMap["бег"])
	require.Equal(t, 13, etalonMap["баскетбол"])
	require.Equal(t, 5, etalonMap["теннис"])
	require.Equal(t, 1, etalonMap["волейбол"])
	require.Equal(t, 0, etalonMap["триатлон"])
	require.Equal(t, 0, etalonMap["трейлранинг"])
	require.Equal(t, 1379, etalonMap["политика"])
	require.Equal(t, 1533, etalonMap["юмор"])

	errorCounts := 0
	for k, v := range etalonMap {
		if int(sk.Count(k)) != v {
			errorCounts++
		}
	}
	fmt.Printf("Accuracy %f \n", 1-float64(errorCounts)/float64(len(etalonMap)))
	return errorCounts
}
func TestCountMinSketchOnYoutubeData(t *testing.T) {
	runCountMinSketchTestsYoutubeData(t, countMinSketchOnYoutubeData, "countMinSketchOnYoutubeData", countMinSketchTestYoutubeData)
}

func simpleMapBench() {
	// https://www.kaggle.com/datasets/datasnaek/youtube-new
	file, err := os.Open("RUvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	etalonMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParse := strings.Split(line, ",")
		tagsLine := lineParse[6]
		tagsLine = strings.ReplaceAll(tagsLine, "\"", "")
		tags := strings.Split(tagsLine, "|")
		for _, tag := range tags {
			etalonMap[tag] += 1
		}
	}
}

func countMinSketchBench(w int, d int) {
	// https://www.kaggle.com/datasets/datasnaek/youtube-new
	file, err := os.Open("RUvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sk := probably.NewSketch(w, d)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineParse := strings.Split(line, ",")
		tagsLine := lineParse[6]
		tagsLine = strings.ReplaceAll(tagsLine, "\"", "")
		tags := strings.Split(tagsLine, "|")
		for _, tag := range tags {
			sk.Increment(tag)
		}
	}
}

func BenchmarkSimpleMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleMapBench()
	}
}

func BenchmarkCountMinSketch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countMinSketchBench(10, 5)
	}
}
