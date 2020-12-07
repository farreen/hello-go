package main

import (
		"fmt"
		"bufio"
		 "os"
       )
func main(){
	words,err:=scanwords("./input.txt")
		if err!=nil{
			panic(err)
		}
	fmt.Println("Printing word by word in order")
	for _, word := range words{
		fmt.Println(word)
	}
	// 0        1    2  3     4
	// nasreen farah is name  my
	// my       name is farah nasreen
	//           i       j       
	fmt.Println("Printing word by word in reverse order")
	   for i, j := 0, len(words)-1; i < j; i, j = i + 1, j-1 {
			words[i], words[j] =  words[j], words[i]
	}
	for _, word := range words{
		fmt.Println(word)
	}
}

func scanwords(path string) ([]string, error){
	file, err := os.Open(path)
		if err!=nil{
			return nil, err
		}
	defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		var words[]string
		for scanner.Scan(){
			words=append(words,scanner.Text())
		}
       return words,nil
}
