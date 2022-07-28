package main

import "fmt"

/*
归并排序
归并排序（Merge Sort）是建立在归并操作上的一种排序算法。它和快速排序一样，采用了分治法。

基本思想
归并的含义是将两个或两个以上的有序表组合成一个新的有序表。也就是说，从几个数据段中逐个选出最小的元素移入新数据段的末尾，使之有序。

那么归并排序的算法我们可以这样理解：

假如初始序列含有 n 个记录，则可以看成是 n 个有序的子序列，每个子序列的长度为 1。然后两两归并，得到 n/2 个长度为2或1的有序子序列；再两两归并，
……，如此重复，直到得到一个长度为 n 的有序序列为止，这种排序方法称为 二路归并排序，下文介绍的也是这种排序方式。

*/


func mergeSort(s []int) []int {
	len := len(s)
	if len == 1 {
		return s //最后切割只剩下一个元素
	}
	m := len / 2
	leftS := mergeSort(s[:m])
	rightS := mergeSort(s[m:])
	return merge(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func merge(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

func main() {

	res := []int{5, 1, 7, 2, 8, 3, 4, 9, 0, 6}
	fmt.Println("排序前：res = ", res)
	result := mergeSort(res)
	fmt.Println("排序后：res = ", result)
}



//func merge(arr []int, L int, M int, R int) {
//	leftSize := M - L + 1
//	rightSize := R - M
//	leftArr := make([]int, leftSize)
//	rightArr := make([]int, rightSize)
//	for i := L; i <= M; i++ {
//		leftArr[i-L] = arr[i]
//	}
//	for j := M + 1; j <= R; j++ {
//		rightArr[j-M-1] = arr[j]
//	}
//	var i, j, k int
//	for {
//		if i < leftSize && j < rightSize {
//			if leftArr[i] < rightArr[j] {
//				arr[k] = leftArr[i]
//			} else {
//				arr[k] = rightArr[j]
//			}
//			k++
//			i++
//			j++
//		}
//	}
//	for {
//		if i < leftSize {
//			arr[k] = leftArr[i]
//			k++
//			i++
//		}
//	}
//
//	for {
//		if j < rightSize {
//			arr[k] = rightArr[i]
//			k++
//			j++
//		}
//	}
//}
//
//func mergeSort(arr []int, L int, R int){
//	if L == R return
//	M := (L + R) / 2
//	mergeSort(arr, L, M)
//	mergeSort(arr, M+1, R)
//	merge(arr, L, M, R)
//}
//
//func main() {
//	res := []int{5, 1, 7, 2, 8, 3, 4, 9, 0, 6}
//	fmt.Println("排序前：res = ", res)
//	mergeSort(res, 0, len(res))
//	fmt.Println("排序后：res = ", res)
//}
