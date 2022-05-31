package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
要求：随机生成五个数，并将其反转打印 , 复杂应用.
*/

func main() {
	var arr [5]int
	length := len(arr)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("交换前 arr = ", arr)
	temp := 0
	for i := 0; i < length/2; i++ {
		temp = arr[length-i-1]
		arr[length-i-1] = arr[i]
		arr[i] = temp
	}
	fmt.Println("交换前 arr = ", arr)

}
