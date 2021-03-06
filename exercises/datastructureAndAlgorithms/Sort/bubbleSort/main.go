package main

import "fmt"

/*

冒泡排序
冒泡排序（Bubble Sort），顾名思义，就是指越小的元素会经由交换慢慢“浮”到数列的顶端。

算法原理
从左到右，依次比较相邻的元素大小，更大的元素交换到右边；
从第一组相邻元素比较到最后一组相邻元素，这一步结束最后一个元素必然是参与比较的元素中最大的元素；
按照大的居右原则，重新从左到后比较，前一轮中得到的最后一个元素不参与比较，得出新一轮的最大元素；
按照上述规则，每一轮结束会减少一个元素参与比较，直到没有任何一组元素需要比较。
*/

func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

//优化后的冒泡排序
func bubbleSortQuicker(arr *[]int) {
	var flag bool
	for i := 0; i < len(*arr)-1; i++ {
		flag = false
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

func main() {
	res := []int{1, 8, 3, 6, 9, 12}
	fmt.Println("排序前：res = ", res)
	bubbleSort(&res)
	fmt.Println("排序前：res = ", res)
	fmt.Println()

	res2 := []int{1, 7, 3, 6, 9, 12}
	fmt.Println("排序前：res = ", res2)
	bubbleSortQuicker(&res2)
	fmt.Println("排序前：res = ", res2)
}
