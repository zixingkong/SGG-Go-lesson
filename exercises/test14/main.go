package main

import "fmt"

/*
说明：编写一个函数 fbn(n int) ，要求完成
1) 可以接收一个 n int
2) 能够将斐波那契的数列放到切片中
3) 提示, 斐波那契的数列形式: arr[0] = 1; arr[1] = 1; arr[2]=2; arr[3] = 3; arr[4]=5; arr[5]=8
*/

func fbn(n int) []uint64 {
	arr := make([]uint64, n)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func main() {
	res := fbn(10)
	fmt.Println(res)
}
