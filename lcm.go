package main

import "fmt"

// func main() {
//
// 	var a, b, l, max int
// 	fmt.Println("Enter two numbers: ")
// 	fmt.Scanf("%d %d", &a, &b)
// 	for l = max; l <= a*b; l++ {
// 		if l%a == 0 && l%b == 0 {
// 			break
// 		}
// 	}
// 	fmt.Printf("lcm = %d\n", l)
//
// }

func giveslcm(a, b int) int {
	var l int

	for l = 1; l <= a*b; l++ {
		if l%a == 0 && l%b == 0 {
			return l
		}
	}
	return l
}

func main() {
	var arr []int = []int{20, 30}
	length := len(arr)
	result := arr[0]
	for i := 1; i < length; i++ {
		result = giveslcm(result, arr[i])
	}
	fmt.Println(result)
}
