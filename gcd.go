package main

import "fmt"

func main() {
	fmt.Println(gcd(10, 5))
}

func gcd(a, b int) int {
	var x, y int
	if a > b {
		x = a
		y = b
	} else {
		x = b
		y = a
	}
	for x%y != 0 {
		temp := x
		x = y
		y = temp % x
	}
	return y
}
