package main

import "fmt"

/*
	编写方法，给指定的二维数组转置
 */

type MethodUtils struct {
}



func (me MethodUtils)Zhuanzhi(arr *[3][3]int) {
	var temp int
	for i:=0;i<len(*arr);i++{
		for j:=0;j<i;j++{
			temp = (*arr)[i][j]
			(*arr)[i][j] = (*arr)[j][i]
			(*arr)[j][i] = temp
		}
	}
}

func main() {
	var me MethodUtils
	arr := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println("arr = ",arr)
	me.Zhuanzhi(&arr)
	fmt.Println("arr = ",arr)
}
