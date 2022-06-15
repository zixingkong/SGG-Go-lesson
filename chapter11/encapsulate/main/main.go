package main

import (
	"SGG-Go-lesson/chapter11/encapsulate/model"
	"fmt"
)

func main() {

	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())

}
