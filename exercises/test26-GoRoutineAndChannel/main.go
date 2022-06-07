package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

/*
1）创建一个Person结构体[Name,Age,Address]
2) 使用rand方法配合随机创建10个Person实例，并放入到channel中
3）遍历channel,将各个Person实例的信息显示在终端
*/

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var personChan chan *Person
	personChan = make(chan *Person, 10)
	name := ""
	age := 0
	address := ""
	for i := 0; i <= 9; i++ {
		age = rand.Intn(1000)
		name = "test_" + strconv.Itoa(age)
		address = "addr_" + strconv.Itoa(age)
		person := Person{
			Name:    name,
			Age:     age,
			Address: address,
		}
		personChan <- &person
	}
	close(personChan)
	for v := range personChan {
		fmt.Println("person = ", *v)
	}
}
