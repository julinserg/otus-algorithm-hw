package main

import (
	"flag"
	"fmt"
	"os"

	p21rle "github.com/julinserg/otus_algorithm_home_work/21-rle"
)

var (
	fromPath, toPath, typeAction string
)

func init() {
	flag.StringVar(&fromPath, "from", "", "file src")
	flag.StringVar(&toPath, "to", "", "file dst")
	flag.StringVar(&typeAction, "type", "encode", "encode/decode")
}

func main() {
	flag.Parse()

	b, err := os.ReadFile(fromPath) // just pass the file name
	if err != nil {
		panic(err)
	}

	var resultRLE string
	if typeAction == "encode" {
		resultRLE = p21rle.RLEEncode(string(b))
	} else if typeAction == "decode" {
		resultRLE = p21rle.RLEDecode(string(b))
	} else {
		fmt.Print("type field should be encode or decode")
	}

	destination, err := os.Create(toPath)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	_, err = destination.Write([]byte(resultRLE))
	if err != nil {
		panic(err)
	}
}
