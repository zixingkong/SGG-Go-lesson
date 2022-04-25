package main

import "fmt"

func fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fn(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*fn(n-1) + 1
	}
}

func peach(n int)int{
	if (n==10){
		return 1
	}else {
		return 2* (peach(n+1)+1)
	}
}


func main() {
	var num int = 6
	fmt.Println(fib(num))

	var num2 int = 5
	fmt.Println(fn(num2))

	num3 := 1
	fmt.Println(peach(num3))
}
