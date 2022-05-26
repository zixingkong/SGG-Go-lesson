package main

import "fmt"

/*
请求出一个数组的最大值，并得到对应的下标。
*/

func main() {
	arr := [...]int{1, 3, 2, 5, 4, 8, 6}
	min := arr[0]
	minIndex := 0
	for index, value := range arr {
		if min < value {
			min = value
			minIndex = index
		}
	}
	fmt.Printf("maxValue = %v, maxValueIndex = %v", min, minIndex)
}
