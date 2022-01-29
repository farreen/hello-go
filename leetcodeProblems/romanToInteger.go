package main

import "fmt"

func main() {
	romanNum := "MCMXCIV"
	fmt.Println(romanToInt(romanNum))
}
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int = 0
	for i, _ := range s {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}
	return sum
}
