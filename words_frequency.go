package main

import (
		"fmt"
		"bufio"
		 "os"
       )
func main(){
	words,err:=scanwords("./words_input.txt")
		if err!=nil{
			panic(err)
		}

                  counts := make(map[string]int)
		for _,word := range words{
			_,ok := counts[word]
				if ok {
					counts[word] +=1
				} else {
					counts[word] =1
				}

		}
	for word := range counts{
		fmt.Println(word,"=>", counts[word])
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
