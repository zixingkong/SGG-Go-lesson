package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("CircleQueue is full")
	}
	cq.array[cq.tail] = val
	cq.tail = (cq.tail + 1) % cq.maxSize
	return
}

func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return -1, errors.New("CircleQueue is empty")
	}
	val = cq.array[cq.head]
	cq.head = (cq.head + 1) % cq.maxSize
	return val, nil
}


func (cq *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	size := cq.Size()
	if size == 0 {
		fmt.Println("环形队列为空")
	}
	tempHead := cq.head
	for i:=0; i<size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, cq.array[tempHead])
		tempHead = (tempHead+1)%cq.maxSize
	}
	fmt.Println()
}



func (cq *CircleQueue) IsFull() bool {
	return (cq.tail+1)%cq.maxSize == cq.head
}

func (cq *CircleQueue) IsEmpty() bool {
	return cq.head == cq.tail
}

func (cq *CircleQueue) Size() int {
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}


func main() {

	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
