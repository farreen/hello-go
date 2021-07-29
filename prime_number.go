package main

import (
	"fmt"
	"math"
)

// find all prime numbers between two given numbers
func main() {
	//	var l, u, i int
	//	fmt.Println("Enter two numbers ")
	//	fmt.Scanf("%d%d", &l, &u)
	//	for x := l + 1; x <= u-1; x++ {
	//		for i = 2; i < x; i++ {
	//			if x%i == 0 {
	//				break
	//			}
	//		}
	//		if x == i {
	//			fmt.Printf("%d ", x)
	//		}
	//	}

	fmt.Println(givesPrime(5, 17))
	//largestPrimefactor(13195)
}

// print all the prime number between two given numbers using function
func givesPrime(a, b int) []int {
	var primes []int
	var i, j int
	for i = a + 1; i <= b-1; i++ {
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == i {
			primes = append(primes, i)
		}
	}
	return primes
}

// Check whether a number is prime or not
func isPrime(n int) bool {
	var i int
	for i = 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Print largest prime factor
func largestPrimefactor(n int) {
	var primes []int
	var max int = 0
	var i int
	for i = 2; i < n; i++ {
		if n%i == 0 && isPrime(i) {
			primes = append(primes, i)
			max = i
		}
	}

	fmt.Println(primes)
	fmt.Println(max)
}
