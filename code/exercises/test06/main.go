package main

import "fmt"

/*
函数可以没有返回值案例，编写一个函数,从终端输入一个整数打印出对应的金子塔
    *
   ***
  *****
*********


 */

func printJinZiTa(num int){
	for i:=1;i<=num;i++{
		for k:=1;k<=num-i;k++{
			fmt.Printf(" ")
		}
		for j:=1;j<=2*i-1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {

	var n int
	fmt.Println("请输入打印金字塔的层数：")
	fmt.Scanln(&n)
	printJinZiTa(n)
	//printJinZiTa(20)
}