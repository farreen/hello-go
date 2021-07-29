package main

import "fmt"

// find all prime numbers between two given numbers
func main() {
	var l, u, i int
	fmt.Println("Enter two numbers ")
	fmt.Scanf("%d%d", &l, &u)
	for x := l + 1; x <= u-1; x++ {
		for i = 2; i < x; i++ {
			if x%i == 0 {
				break
			}
		}
		if x == i {
			fmt.Printf("%d ", x)
		}
	}
}
