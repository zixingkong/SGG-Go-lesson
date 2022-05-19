package main

import "fmt"

/*
面试题：冒泡排序算法实现
 */

func bubbleSort(arr *[]int){
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-i-1;j++{
			if (*arr)[j] > (*arr)[j+1]{
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}

		}
	}
}



func main() {
	res := []int{1,8,3,6,9,12}
	fmt.Println("排序前：res = ",res)
	bubbleSort(&res)
	fmt.Println("排序前：res = ",res)

}