package main

import (
	"flag"
	"fmt"
	"os"

	p21rle "github.com/julinserg/otus_algorithm_home_work/21-rle"
)

var (
	fromPath, toPath, mode string
)

func init() {
	flag.StringVar(&fromPath, "from", "", "file src")
	flag.StringVar(&toPath, "to", "", "file dst")
	flag.StringVar(&mode, "mode", "encode", "encode/decode")
}

func main() {
	flag.Parse()

	b, err := os.ReadFile(fromPath) // just pass the file name
	if err != nil {
		panic(err)
	}

	var resultRLE string
	if mode == "encode" {
		resultRLE = p21rle.RLEEncode(string(b))
	} else if mode == "decode" {
		resultRLE = p21rle.RLEDecode(string(b))
	} else {
		fmt.Print("type field must be encode or decode")
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
