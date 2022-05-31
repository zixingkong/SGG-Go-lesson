package main

import "fmt"

type Student struct {
	name string
	num  int
}

func TypeJudge(items ...interface{}) {
	for index, item := range items {
		switch item.(type) {

		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, item)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, item)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, item)
		case int32:
			fmt.Printf("第%v个参数是 int32 类型，值是%v\n", index, item)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, item)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, item)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, item)
		default:
			fmt.Printf("第%v个参数类型不确定，值是%v\n", index, item)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	var student1 Student = Student{
		name: "小白",
		num:  1,
	}
	var student2 *Student = &Student{
		name: "小黑",
		num:  2,
	}
	TypeJudge(n1, n2, n3, n4, name, address, student1, student2)
}
