package main

import "fmt"

func powerIterate(num int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * num
	}
	return res
}

func fibonacciIterate(index int) int {
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	a := 0
	b := 1
	for i := 0; i < index-1; i++ {
		t := b
		b = b + a
		a = t
	}
	return b
}

func fibonacciRecursive(index int) int {
	if index <= 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return fibonacciRecursive(index-1) + fibonacciRecursive(index-2)
}

func primeNumbersBruteforce(limit int) int {
	count := 0
	for i := 2; i <= limit; i++ {
		isPrime := true
		for d := 2; d <= limit; d++ {
			if i%d == 0 && i != d {
				isPrime = false
			}
		}
		if isPrime {
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("powerIterate num: %d, power: %d. Result: %d \n", 2, 4, powerIterate(2, 4))
	fmt.Printf("fibonacciIterate index: %d. Result: %d \n", 10, fibonacciIterate(10))
	fmt.Printf("fibonacciRecursive index: %d. Result: %d \n", 10, fibonacciRecursive(10))
	fmt.Printf("primeNumbersBruteforce limit: %d. Result: %d \n", 100, primeNumbersBruteforce(100))
}
