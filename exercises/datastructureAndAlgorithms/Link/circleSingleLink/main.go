package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode.no, newCatNode.name, "加入到环形列表")
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表")
		return
	}
	for {
		fmt.Printf("猫的信息为=[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("这是一个空的单向环形列表，不能删除")
		return head
	}
	// 如果只有一个节点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		}
		return head
	}
	//将heLper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 如果包含两个及两个以上的节点
	flag := true
	for {
		if temp.next == head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.next.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	//这里还比较一次
	if flag {
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}

	}
	return head
}

func main() {

	//这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 30)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)
}
