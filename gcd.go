package main

import "fmt"

func main() {
	fmt.Println(gcd(10, 5))
}

func gcd(a, b int) int {
	var max, min int
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	for max%min != 0 {
		temp := max
		max = min
		min = temp % max
	}
	return min
}
