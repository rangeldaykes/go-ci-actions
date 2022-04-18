package main

func Fibonacci(nSeries int) int {

	if nSeries <= 1 {
		return nSeries
	}

	return Fibonacci(nSeries-1) + Fibonacci(nSeries-2)
}
