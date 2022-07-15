package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		}
		if temp.next.no > newHeroNode.no {
			break
		}
		if temp.next.no == newHeroNode.no {
			flag = true
			break
		}

		temp = temp.next
	}

	if flag {
		fmt.Println("SingleLink has exist HeroNode no")
		return
	}

	newHeroNode.next = temp.next
	temp.next = newHeroNode

}

func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := false

	for {
		if temp.next == nil {
			break
		}
		if temp.next.no == id {
			flag = true
			break
		}

		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
	}else{
		fmt.Println("sorry, 要删除的id不存在")
	}

}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil{
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil{
			break
		}
	}
}

func main() {

	//1. 先创建一个头结点,
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no : 1,
		name : "宋江",
		nickname : "及时雨",
	}

	hero2 := &HeroNode{
		no : 2,
		name : "卢俊义",
		nickname : "玉麒麟",
	}

	hero3 := &HeroNode{
		no : 3,
		name : "林冲",
		nickname : "豹子头",
	}

	// hero4 := &HeroNode{
	// 	no : 3,
	// 	name : "吴用",
	// 	nickname : "智多星",
	// }

	//3. 加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)

	// 4. 显示
	ListHeroNode(head)

	// 5 删除
	fmt.Println()
	DelHerNode(head, 1)
	DelHerNode(head, 3)
	ListHeroNode(head)
}
