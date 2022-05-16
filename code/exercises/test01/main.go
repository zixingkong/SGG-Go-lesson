package main

import "fmt"

/*
 题 1：斐波那契数 请使用递归的方式，求出斐波那契数 1,1,2,3,5,8,13...
	给你一个整数 n，求出它的斐波那契数是多少？
*/

func fibonacci(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	res := fibonacci(3)
	fmt.Println("res = ", res)
	fmt.Println("res = ", fibonacci(4))
	fmt.Println("res = ", fibonacci(5))
	fmt.Println("res = ", fibonacci(6))

}
