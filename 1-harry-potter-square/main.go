package main

import "fmt"

func main() {
	n := 25
	for x := 0; x < n; x++ {
		fmt.Printf("\n")
		for y := 0; y < n; y++ {
			if y < 2*x+2 && y > 2*x-1 {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}

		}
	}
}
