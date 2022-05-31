package main

import "fmt"

/*
编写一个方法算该矩形的面积(可以接收长 len，和宽 width)，
将其作为方法返回值。在 main 方法中调用该方法，接收返回的面积值并打印。
*/
type Rectangle struct {
	Len   float64
	Width float64
}

func (re Rectangle) area() float64 {
	return re.Len * re.Width
}

/*
编写方法：判断一个数是奇数还是偶数
*/
func (re Rectangle) Judge(num int) {
	if num%2 == 0 {
		fmt.Printf("%v 是偶数\n", num)
	} else {
		fmt.Printf("%v 是奇数\n", num)
	}
}

/*
) 根据行、列、字符打印 对应行数和列数的字符，比如：行：3，列：2，字符*,则打印相应的效 果
*/
func (re Rectangle) printTest(m, n int, flag string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(flag)
		}
		fmt.Println()
	}
}
func main() {
	re := Rectangle{5, 3}
	fmt.Println("area = ", re.area())
	re.Judge(5)
	re.Judge(6)
	re.printTest(5, 6, "+")
}
