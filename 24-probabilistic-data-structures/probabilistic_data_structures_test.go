package p24probabilisticdatastructures

import (
	"bufio"
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

func TestCountMinSketchOnBigData(t *testing.T) {
	// https://www.kaggle.com/datasets/datasnaek/youtube-new
	file, err := os.Open("RUvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sk := probably.NewSketch(1<<20, 5)

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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	require.Equal(t, 367, int(sk.Count("футбол")))
	require.Equal(t, 84, int(sk.Count("хоккей")))
	require.Equal(t, 43, int(sk.Count("бокс")))
	require.Equal(t, 33, int(sk.Count("смешанные единоборства")))
	require.Equal(t, 32, int(sk.Count("плавание")))
	require.Equal(t, 12, int(sk.Count("бег")))
	require.Equal(t, 13, int(sk.Count("баскетбол")))
	require.Equal(t, 5, int(sk.Count("теннис")))
	require.Equal(t, 1, int(sk.Count("волейбол")))
	require.Equal(t, 0, int(sk.Count("триатлон")))
	require.Equal(t, 0, int(sk.Count("трейлранинг")))
	require.Equal(t, 1379, int(sk.Count("политика")))
}
