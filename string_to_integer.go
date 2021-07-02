package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//words, err := scanwords(os.Args[1])

	words, err := scanwords("./numbers.txt")
	if err != nil {
		panic(err)
	}
	for _, word := range words {
		nums, _ := strconv.Atoi(word)
		result := nums * 2
		fmt.Println(nums, "* 2 =", result)
	}
}

func scanwords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}
