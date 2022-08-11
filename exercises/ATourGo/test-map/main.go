package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stringSlice := strings.Split(s," ")
	fmt.Println(stringSlice)
	wordMap := make(map[string]int)
	for _,word := range stringSlice{
		_, ok := wordMap[word]
		if ok {
			wordMap[word] +=1
		}else{
			wordMap[word] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
