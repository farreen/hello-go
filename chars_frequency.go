package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := scanChars("./input.txt")
	if err != nil {
		panic(err)
	}

	counts := make(map[byte]int)
	for _, char := range bytes {
		_, ok := counts[char]
		if ok {
			counts[char] += 1
		} else {
			counts[char] = 1
		}
	}
	for char := range counts {
		fmt.Printf("%c =>  %d \n", char, counts[char])
	}
}

func scanChars(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	return buffer, nil

}
