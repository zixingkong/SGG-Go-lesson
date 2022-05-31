package main

import "fmt"

/*
题 3：猴子吃桃子问题 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，
然后 再多吃一个。当到第十天时，想再吃时（还没吃），发现只有 1 个桃子了。问题：最初共多少个桃子？
*/

/*
 	第10天 1个
	第9天 2*(f(10)+1)

*/

func getPeach(n int) int {
	if n == 10 {
		return 1
	} else {
		return 2 * (getPeach(n+1) + 1)
	}
}

func main() {
	res := getPeach(1)
	fmt.Println("第1天有 ", res)
}
