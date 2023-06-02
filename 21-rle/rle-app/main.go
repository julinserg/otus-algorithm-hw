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

	b, err := os.ReadFile(fromPath)
	if err != nil {
		panic(err)
	}

	var resultRLE []byte
	if mode == "encode" {
		resultRLE = p21rle.RleEncode(b)
	} else if mode == "decode" {
		resultRLE = p21rle.RleDecode(b)
	} else {
		fmt.Print("type field must be encode or decode")
	}

	destination, err := os.Create(toPath)
	if err != nil {
		panic(err)
	}
	defer destination.Close()

	_, err = destination.Write(resultRLE)
	if err != nil {
		panic(err)
	}
}
