package main

import "fmt"

/*
6) 定义小小计算器结构体(Calcuator)，
实现加减乘除四个功能 实现形式
1：分四个方法完成: 实现形式
2：用一个方法搞定
*/
type Calculator struct {
	Num1 float64
	Num2 float64
}

func (ca Calculator) Calculate(operator string) float64 {
	var res float64
	switch operator {
	case "+":
		res = ca.Num1 + ca.Num2
	case "-":
		res = ca.Num1 - ca.Num2
	case "*":
		res = ca.Num1 * ca.Num2
	case "/":
		res = ca.Num1 / ca.Num2
	default:
		fmt.Println("输入的运算符有误")
	}
	return res
}

func main() {
	ca := Calculator{5, 6.0}
	res1 := ca.Calculate("+")
	res2 := ca.Calculate("-")
	res3 := ca.Calculate("*")
	res4 := ca.Calculate("/")
	fmt.Printf("%v + %v = %v\n", ca.Num1, ca.Num2, res1)
	fmt.Printf("%v - %v = %v\n", ca.Num1, ca.Num2, res2)
	fmt.Printf("%v * %v = %v\n", ca.Num1, ca.Num2, res3)
	fmt.Printf("%v / %v = %v\n", ca.Num1, ca.Num2, res4)
}
