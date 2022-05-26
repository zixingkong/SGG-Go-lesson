package main

import "fmt"

/*
1) 创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。
使用 for 循环访问所有元素并打印 出来。提示：字符数据运算 'A'+1 -> 'B
*/

func main() {
	var arr [26]byte = [26]byte{}
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}

	for _,value := range arr{
		fmt.Printf("%c ",value)
	}
}
