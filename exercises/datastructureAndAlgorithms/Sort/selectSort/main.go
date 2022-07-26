package main

import "fmt"

/*
选择排序
选择排序（Selection Sort）是一种简单直观的排序算法。它的基本思想就是，每一趟 n-i+1(i=1,2,...,n-1)个记录中选取关键字最小的记录作为有序序列的第 i 个记录。

算法步骤
简单选择排序：

在未排序序列中找到最小（大）元素，存放到排序序列的起始位置;
在剩余未排序元素中继续寻找最小（大）元素，放到已排序序列的末尾;
重复步骤2，直到所有元素排序完毕。

算法分析
选择排序是不稳定排序，时间复杂度固定为 O(n²)，因此它不适用于数据规模较大的序列。不过它也有优点，就是不占用额外的内存空间。
*/

func swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func selectSort(arr *[]int) {
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if (*arr)[j] < (*arr)[min] {
				min = j
			}
		}
		swap(*arr, i, min)
	}
}

func main() {
	res := []int{5, 1, 7, 2, 8, 3, 4, 9, 0, 6}
	fmt.Println("排序前：res = ", res)
	selectSort(&res)
	fmt.Println("排序后：res = ", res)
}
