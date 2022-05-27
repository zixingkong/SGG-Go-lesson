package main

import (
	"SGG-Go-lesson/exercises/test23/module"
	"fmt"
)

func main() {

	stu := module.NewStudent("小白", 85)
	fmt.Println("stu = ", *stu)
	fmt.Printf("name = %v score= %v", stu.Name, stu.GetScore())

}
