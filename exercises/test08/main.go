package main

import "fmt"

/*
 编写函数,对给定的一个二维数组(3×3)转置，这个题讲数组的时候再完成
*/

func ZhuanZhi(arr *[3][3]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < i;j++ {
			temp := (*arr)[i][j]
			(*arr)[i][j] = (*arr)[j][i]
			(*arr)[j][i] = temp
		}
	}

}

func main() {
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ZhuanZhi(&arr)
	fmt.Println(arr)
}
