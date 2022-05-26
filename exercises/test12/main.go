package main

import "fmt"

/*
请求出一个数组的和和平均值
*/

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12,11}
	sum := 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("数组的和为%v,数组的平均值为%v", sum, sum/len(arr))
}
