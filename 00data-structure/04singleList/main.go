package main

import "fmt"

//
type HeroNode struct {
	num int
	name string
	nickname string
	next *HeroNode
}

//InsertTail 在链表尾部插入节点
func InsertTail(new *HeroNode,head *HeroNode)  {
	//先找到链表最后的节点
	tempHead := head
	for {
		if tempHead.next == nil {
			break
		}
		tempHead = tempHead.next
	}
	fmt.Println("add..")
	tempHead.next = new
}
//InsertHeroByNum 按照编号顺序有序添加到链表中
func InsertHeroByNum(new, head *HeroNode) {
	temp := head
	//让插入的节点的num和temp的下一个节点的num比较
	for  {
		if temp.next==nil{
			temp.next = new
			return
		}
		if temp.next.num >= new.num{
			new.next = temp.next
			temp.next = new
			break
		}
		temp = temp.next
	}
}
//DeleteNode 按照英雄名字删除节点
func DeleteNode(head *HeroNode, name string)  {
	if head.next == nil {
		fmt.Println("链表为空。。。")
		return
	}
	temp := head
	for  {
		if temp.next.name == name {
			temp.next = temp.next.next
			break
		}else if temp.next == nil{
			fmt.Println("没有该节点。。。")
			break
		}
		temp = temp.next
	}
}

//ShowHeroList 显示链表
func ShowHeroList(head *HeroNode) {
	if head.next == nil {
		fmt.Println("the list is nil")
		return
	}
	//1.创建一个辅助节点
	temp := head.next
	//2.遍历链表
	for  {
		fmt.Printf("[%d, %s, %s]\t", temp.num,temp.name,temp.nickname)
		temp = temp.next
		if temp == nil {
			break
		}
	}
	fmt.Println()
}
func main() {
	//1.先创建一个头节点
	head := &HeroNode{}
	//2.创建一个新的节点
	hero1 := &HeroNode{
		num: 1,
		name: "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		num: 2,
		name: "武松",
		nickname: "大老虎",
	}

	hero3 := &HeroNode{
		num: 3,
		name: "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		num: 4,
		name: "孙二娘",
		nickname: "黑店",
	}
	InsertHeroByNum(hero1,head)
	InsertHeroByNum(hero3,head)
	InsertHeroByNum(hero2,head)
	InsertHeroByNum(hero4,head)
	ShowHeroList(head)
	DeleteNode(head, "林冲")
	ShowHeroList(head)
	return
}
