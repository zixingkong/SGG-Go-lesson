package main

import "fmt"

/*
堆排序
堆排序（Heap Sort）是指利用堆这种数据结构所设计的一种排序算法。堆的特点：

一颗完全二叉树（也就是会所生成节点的顺序是：从上往下、从左往右）
每一个节点必须满足父节点的值不大于/不小于子节点的值
基本思想
	实现堆排序需要解决两个问题：

		如何将一个无序序列构建成堆？

		如何在输出堆顶元素后，调整剩余元素成为一个新的堆？

以升序为例，算法实现的思路为：

1. 建立一个 build_heap 函数，将数组 tree[0,...n-1] 建立成堆，n 表示数组长度。函数里需要维护的是所有节点的父节点，最后一个子节点下标为 n-1，那么它对应的父节点下标就是(n-1-1)/2。
2. 构建完一次堆后，最大元素就会被存放在根节点 tree[0]。将 tree[0] 与最后一个元素交换，每一轮通过这种不断将最大元素后移的方式，来实现排序。
3. 而交换后新的根节点可能不满足堆的特点了，因此需要一个调整函数 heapify 来对剩余的数组元素进行最大堆性质的维护。
   如果 tree[i] 表示其中的某个节点，那么 tree[2*i+1] 是左孩子，tree[2*i+2] 是右孩子，选出三者中的最大元素的下标，存放于 max 值中，若 max 不等于 i，则将最大元素交换到 i 下标的位置。
   但是，此时以 tree[max] 为根节点的子树可能不满足堆的性质，需要递归调用自身。
*/

func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}

//堆排序
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}

// ---------------------------------------------------------------------------------------------------------------------
// 本例为最小堆
// 最大堆只需要修改less函数即可
type Heap []int

func (h Heap) Init() {
	n := len(h)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)
	return x
}

func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h) {
			break // i已经是叶子结点了
		}
		j := l
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r // 右孩子
		}
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}
