package main

import "fmt"

/*
编写一个函数,从终端输入一个整数(1—9),打印出对应的乘法表
 1*1=1

 */
func printChengfa(num int){
	for i:=1;i<=num;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%v * %v = %v \t",j,i,j*i )
		}
		fmt.Println()
	}

}


func main() {
	var n int
	fmt.Println("请输入乘法表的层数")
	fmt.Scanln(&n)
	printChengfa(n)
}