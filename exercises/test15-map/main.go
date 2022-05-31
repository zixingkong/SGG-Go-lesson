package main

import "fmt"

/*
课堂练习：演示一个 key-value 的 value 是 map 的案例
比如：我们要存放 3 个学生信息, 每个学生有 name 和 sex 信息
思路: map[string]map[string]string
*/

func main() {
	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "tom"
	stuMap["stu01"]["sex"] = "male"
	stuMap["stu02"] = make(map[string]string, 2)
	stuMap["stu02"]["name"] = "mary"
	stuMap["stu02"]["sex"] = "female"
	fmt.Println("stuMap: ", stuMap)
}
