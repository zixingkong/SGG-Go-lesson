package main

import (
	"SGG-Go-lesson/exercises/test24-encapsulation/model"
	"fmt"
)

func main() {
	//创建一个 account 变量
	account := model.NewAccount("11111asfd", "000000", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
	account.Deposite(100,"000000")
	account.GetBalance("000000")
	account.WithDraw(50,"000000")
	account.GetBalance("000000")
	account.SetBalance("000000",200)
	account.GetBalance("000000")
}
