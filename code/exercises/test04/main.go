package main

import "fmt"

/*
请编写一个函数 swap(n1 *int, n2 *int) 可以交换 n1 和 n2 的值
*/
func swap(num1 *int, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

func main() {
	a := 3
	b := 4
	fmt.Printf("a = %v,b = %v\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %v,b = %v", a, b)

}
