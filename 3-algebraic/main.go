package main

import "fmt"

func powerIterate(num int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * num
	}
	return res
}

func powerRecursive(num int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return num
	}
	return num * powerRecursive(num, power-1)
}

func powerFast(num int, power int) int {
	powerL := power
	d := num
	p := 1
	for powerL >= 1 {
		if powerL%2 != 0 {
			p = p * d
		}
		powerL = powerL / 2
		d = d * d
	}
	return p
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

func matrixMultipl(matrixA [2][2]int, matrixB [2][2]int) [2][2]int {
	var result [2][2]int
	result[0][0] = matrixA[0][0]*matrixB[0][0] + matrixA[0][1]*matrixB[1][0]
	result[0][1] = matrixA[0][0]*matrixB[0][1] + matrixA[0][1]*matrixB[1][1]
	result[1][0] = matrixA[1][0]*matrixB[0][0] + matrixA[1][1]*matrixB[1][0]
	result[1][1] = matrixA[1][0]*matrixB[0][1] + matrixA[1][1]*matrixB[1][1]
	return result
}

func fibonacciMatrix(index int) int {
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	a := [2][2]int{
		{1, 1},
		{1, 0},
	}

	powerL := index - 1
	d := a
	p := [2][2]int{
		{1, 0},
		{0, 1},
	}
	for powerL >= 1 {
		if powerL%2 != 0 {
			p = matrixMultipl(p, d)
		}
		powerL = powerL / 2
		d = matrixMultipl(d, d)
	}
	return p[0][0]
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

	fmt.Printf("powerRecursive num: %d, power: %d. Result: %d \n", 2, 5, powerRecursive(2, 5))
	fmt.Printf("powerFast num: %d, power: %d. Result: %d \n", 2, 5, powerFast(2, 5))
	fmt.Printf("fibonacciMatrix index: %d. Result: %d \n", 10, fibonacciMatrix(10))
}
