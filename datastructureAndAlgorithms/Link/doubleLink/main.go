package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
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
	newHeroNode.pre = temp

}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		} else if temp.next.no > newHeroNode.no {
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("sorry, no=%d has exist", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil{
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

func DelHeroNode(head *HeroNode,id int){
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		}else if temp.next.no == id {
			flag = true
			break
		}

		temp = temp.next
	}
	if flag{
		temp.next = temp.next.next
		if temp.next != nil{
			temp.next.pre = temp
		}
	}else{
		fmt.Printf("sorry, no=%d has exist", id)
	}
}

func ListHeroNode(head *HeroNode) {
	temp:= head
	if temp.next == nil{
		fmt.Println("empty")
		return
	}
	for{
		fmt.Printf("[%d , %s , %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil{
			break
		}
	}

}

func ListHeroNode2(head *HeroNode) {
	temp:= head
	if temp.next == nil{
		fmt.Println("empty")
	}
	for{
		if temp.next == nil{
			break
		}
		temp= temp.next
	}
	for{
		fmt.Printf("[%d , %s , %s]==>", temp.no,
			temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil{
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


	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
	fmt.Println("逆序打印")
	ListHeroNode2(head)

}
