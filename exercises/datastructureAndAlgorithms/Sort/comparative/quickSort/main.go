package main

import (
	"fmt"
	"math/rand"
)

/*
快速排序
	快速排序（Quick Sort），是冒泡排序的改进版，之所以“快速”，是因为使用了分治法。它也属于交换排序，通过元素之间的位置交换来达到排序的目的。

	基本思想
	在序列中随机挑选一个元素作基准，将小于基准的元素放在基准之前，大于基准的元素放在基准之后，再分别对小数区与大数区进行排序。

	一趟快速排序的具体做法是：

	设两个指针 i 和 j，分别指向序列的头部和尾部；
	先从 j 所指的位置向前搜索，找到第一个比基准小的值，把它与基准交换位置；
	再从 i 所指的位置向后搜索，找到第一个比基准大的值，把它与基准交换位置；
	重复 2、3 两步，直到 i = j。
	仔细研究一下上述算法我们会发现，在排序过程中，对基准的移动其实是多余的，因为只有一趟排序结束时，也就是 i = j 的位置才是基准的最终位置。

	由此可以优化一下算法：

	设两个指针 i 和 j，分别指向序列的头部和尾部；
	先从 j 所指的位置向前搜索，找到第一个比基准小的数值后停下来，再从 i 所指的位置向后搜索，找到第一个比基准大的数值后停下来，把 i 和 j 指向的两个值交换位置；
	重复步骤2，直到 i = j，最后将相遇点指向的值与基准交换位置。


算法分析
	快速排序是不稳定排序，它的平均时间复杂度为 O(nlogn)，平均空间复杂度为 O(logn)。

	快速排序中，基准的选取非常重要，它将影响排序的效率。举个例子，假如序列本身顺序随机，快速排序是所有同数量级时间复杂度的排序算法中平均性能最好的，
	但如果序列本身已经有序或基本有序，直接选取固定位置，例如第一个元素作为基准，会使快速排序就会沦为冒泡排序，时间复杂度为 O(n^2)。为了避免发生这种情况，引入下面两种获取基准的方法：

	随机选取

		就是选取序列中的任意一个数为基准的值。

	三者取中

		就是取起始位置、中间位置、末尾位置指向的元素，对这三个元素排序后取中间数作为基准。

	经验证明，三者取中的规则可以大大改善快速排序在最坏情况下的性能。
*/

// 交换数组两个位置的元素
func swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

// 选取第一个元素作为基准
func selectPivot(arr []int, low int) int {
	return arr[low]
}

// 选取随机元素作为基准: 随机选择基准的位置，区间在 low 和 high 之间
func selectPivot2(arr []int, low int, high int) int {
	pivot := rand.Intn(high-low) + low
	swap(arr, low, pivot)
	return arr[low]
}

// 就是取起始位置、中间位置、末尾位置指向的元素，对这三个元素排序后取中间数作为基准。
func selectPivot3(arr []int, low int, high int) int {
	mid := low + (high-low)/2
	if arr[mid] > arr[high] {
		swap(arr, mid, high)
	}
	if arr[low] > arr[high] {
		swap(arr, low, high)
	}
	if arr[mid] > arr[low] {
		swap(arr, low, mid)
	}
	return arr[low]
}

func quickSort(arr []int, low int, high int) {
	var i, j, pivot int
	if low >= high {
		return
	}
	pivot = selectPivot3(arr, low, high)
	i = low
	j = high
	for {
		if i == j {
			break
		}
		for {
			if arr[j] >= pivot && i < j {
				j--
			} else {
				break
			}
		}

		for {
			if arr[i] <= pivot && i < j {
				i++
			} else {
				break
			}
		}

		if i < j {
			swap(arr, i, j)
		}
	}

	arr[low] = arr[i]
	arr[i] = pivot

	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}

func main() {
	res := []int{5, 1, 7, 2, 8, 3, 4, 9, 0, 6}
	fmt.Println("排序前：res = ", res)
	quickSort(res, 0, len(res)-1)
	fmt.Println("排序后：res = ", res)
}
