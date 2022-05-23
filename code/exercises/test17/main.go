package main

import "fmt"

/*
编写结构体(MethodUtils)，
编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形， 在 main 方法中调用该方法。
*/

type MethodUtils struct {
}

/*
编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
*/
func (m MethodUtils) printJuXing() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) printJuXing2(m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	m := MethodUtils{}
	m.printJuXing()
	m.printJuXing2(5, 3)
}
