package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func main() {
	m := map[string]func(a, b int) int{
		"+": add,
		"-": minus,
		"*": multiply,
		"/": divide,
	}
	var op = "+"
	a, b := 3, 2
	result := m[op](a, b)
	fmt.Println(result)

	/*	for op, operation := range m {
		fmt.Println(op, operation(2, 2))
	}*/
}
