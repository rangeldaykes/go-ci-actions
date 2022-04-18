package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(4))
}

func Fibonacci(nSeries int) int {

	if nSeries <= 1 {
		return nSeries
	}

	return Fibonacci(nSeries-1) + Fibonacci(nSeries-2)
}
