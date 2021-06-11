package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	const maxsize = 100

	buffer := make([]byte, maxsize)

	for {
		Totalread, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("bytes read", Totalread)
		fmt.Println("bytes stream to string", string(buffer[:Totalread]))
	}
}
