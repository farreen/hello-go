package main

import "fmt"

func main() {
	var number, i int
	fmt.Println("Enter a number ")
	fmt.Scanf("%d", &number)

	for i = 2; i <= number-1; i++ {
		if number%i == 0 {
			break
		}
	}

	if number == i {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a Prime number")
	}
}
