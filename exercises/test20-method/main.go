package main

import "fmt"

/*

 */

type MethodUtils struct {
}

func (me MethodUtils)Printest(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	var me MethodUtils
	me.Printest(5)
	fmt.Println()
	me.Printest(9)
}
