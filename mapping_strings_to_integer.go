package main

import "fmt"

func main() {
	celebs := map[string]int{ //mapping strings to integers
		"Nicolas Cage":       50,
		"Selena Gomez":       21,
		"Jude Law":           41,
		"Scarlett Johansson": 29,
	}

	fmt.Printf("%#v", celebs)
}
