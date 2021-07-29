package main

import "fmt"

func main() {
	fmt.Println(giveFactors(10))
}

func giveFactors(number int) []int {
	var factors []int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
