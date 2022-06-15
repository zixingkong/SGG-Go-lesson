package main

import (
	"SGG-Go-lesson/chapter11/encapexercise/model"
	"fmt"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "000111", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}