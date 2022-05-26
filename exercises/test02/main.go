package main

import "fmt"

/*

已知 f(1)=3; f(n) = 2*f(n-1)+1; 请使用递归的思想编程，求出 f(n)的值?
*/

func getNum(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*getNum(n-1) + 1
	}

}

func main() {
	fmt.Println("getNum(1) = ", getNum(1))
	fmt.Println("getNum(1) = ", getNum(5))
}
