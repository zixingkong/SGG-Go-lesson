package main

import (
	"errors"
	"fmt"
	"os"
)

type SingleQueue struct {
	MaxSize int
	Array   [5]int
	front   int
	rear    int
}

func (this *SingleQueue) AddSingleQueue(val int) (err error) {
	if this.rear == this.MaxSize-1 {
		return errors.New("SingleQueue is full")
	}
	this.rear++
	this.Array[this.rear] = val
	return
}

func (this *SingleQueue) GetSingleQueue() (val int, err error) {
	if this.rear == this.front {
		return -1, errors.New("SingleQueue is empty")
	}
	this.front++
	val = this.Array[this.front]
	return val, nil
}

func (this *SingleQueue) ShowSingleQueue() {
	fmt.Println("队列当前的值为：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("Array[%d]=%d\t", i, this.Array[i])
	}
	fmt.Println()
}

func main() {
	singleQueue := SingleQueue{
		MaxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示添加数据到队列")
		fmt.Println("3. 输入show 表示添加数据到队列")
		fmt.Println("4. 输入exit 表示添加数据到队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := singleQueue.AddSingleQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val, err := singleQueue.GetSingleQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出一个数：", val)
			}
		case "show":
			singleQueue.ShowSingleQueue()
		case "exit":
			os.Exit(0)

		}
	}

}
