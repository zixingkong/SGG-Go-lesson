package main

import "fmt"

/*

计数排序
计数排序（Counting Sort）是一种非比较性质的排序算法，利用了桶的思想。它的核心在于将输入的数据值转化为键存储在额外开辟的辅助空间中，也就是说这个辅助空间的长度取决于待排序列中的数据范围。

如何转化成桶思想来理解呢？我们设立 r 个桶，桶的键值分别对应从序列最小值升序到最大值的所有数值。接着，按照键值，依次把元素放进对应的桶中，然后统计出每个桶中分别有多少元素，再通过对桶内数据的计算，即可确定每一个元素最终的位置。

算法步骤
1. 找出待排序列中最大值 max 和最小值 min，算出序列的数据范围 r = max - min + 1，申请辅助空间 C[r]；
2. 遍历待排序列，统计序列中每个值为 i 的元素出现的次数，记录在辅助空间的第 i 位；
3. 对辅助空间内的数据进行计算（从空间中的第一个元素开始，每一项和前一项相加），以确定值为 i 的元素在数组中出现的位置；
4. 反向填充目标数组：将每个元素 i 放在目标数组的第 C[i] 位，每放一个元素就将 C[i] 减1，直到 C 中所有值都是

算法分析
计数排序属于非交换排序，是稳定排序，适合数据范围不显著大于数据数量的序列。

时间复杂度
它的时间复杂度是线性的，为 O(n+r)，r 表示待排序列中的数据范围，也就是桶的个数。可以这样理解：将 n 个数据依次放进对应的桶中，再从 r 个桶中把数据按顺序取出来。

空间复杂度
占用额外内存，还需要 r 个桶，因此空间复杂度是 O(n+r)，计数排序快于任何比较排序算法，但这是通过牺牲空间换取时间来实现的。
*/

func countingSort(arr []int) []int {
	maxValue := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return arr
}

func main() {
	nums := []int{2, 4, 3, 1, 4, 6, 4, 2}
	fmt.Println("排序前：")
	fmt.Println(nums)
	nums = countingSort(nums)
	fmt.Println("排序后：")
	fmt.Println(nums)
}
