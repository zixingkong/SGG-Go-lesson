package model

import "fmt"

/*
要求
1) 创建程序,在 model 包中定义 Account 结构体： 在 main 函数中体会 Golang 的封装性。
2) Account 结构体要求具有字段： 账号（长度在 6-10 之间） 、 余额(必须>20)、 密码（必须是六
3) 通过 SetXxx 的方法给 Account 的字段赋值。 (同学们自己完成)
4) 在 main 函数中测试
*/

type Account struct {
	accountNum string
	pwd        string
	balance    float64
}

func NewAccount(accountNum string, pwd string, balance float64) *Account {
	if len(accountNum) < 6 || len(accountNum) > 10 {
		fmt.Println("输入的账号不合法")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("输入的密码不合法")
		return nil
	}
	if balance <= 20 {
		fmt.Println("输入的余额不合法")
		return nil
	}
	account := &Account{
		accountNum: accountNum,
		pwd:        pwd,
		balance:    balance,
	}
	return account
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("请输入正确的金额")
		return
	}
	account.balance += money
	fmt.Println("存款成功")
}

func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	if money <= 0 || money > account.balance {
		fmt.Println("请输入正确的金额")
		return
	}
	account.balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	fmt.Println("账号的余额为", account.balance)
}

func (account *Account) SetBalance(pwd string, balance float64) {
	if pwd != account.pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	account.balance = balance

}

func (account *Account) GetBalance(pwd string) {
	if pwd != account.pwd {
		fmt.Println("输入的密码不正确")
		return
	}
	fmt.Println("余额为", account.balance)

}
