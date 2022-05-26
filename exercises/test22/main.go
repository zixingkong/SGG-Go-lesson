package main

import "fmt"

/*
1) 编写一个 Dog 结构体， 包含 name、 age、 weight 字段
2) 结构体中声明一个 say 方法， 返回 string 类型， 方法返回信息中包含所有字段值。
3) 在 main 方法中， 创建 Dog 结构体实例(变量)， 并访问 say 方法， 将调用结果打印输出。
*/

type Dog struct {
	name   string
	age    int
	weight float64
}

func (dog *Dog) say() string {
	info := fmt.Sprintf("dog的信息 name=[%v] age=[%v] weight=[%v]",dog.name ,dog.age,dog.weight)
	return info
}

func main() {
	dog := &Dog{
		name: "小白",
		age: 6,
		weight: 20,
	}
	res := dog.say()
	fmt.Println(res)
}
