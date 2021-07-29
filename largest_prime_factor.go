package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPrimefactor(13195))
}

func largestPrimefactor(number int) int {
	var max int = 0
	for i := 2; i < number; i++ {
		if number%i == 0 && isPrime(i) {
			max = i
		}
	}
	return max
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
