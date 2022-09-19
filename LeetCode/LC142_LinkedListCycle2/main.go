package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func initNode(val int)*ListNode{
	return &ListNode{Val: val}
}

func printListNode(l *ListNode)  {
	for l!=nil{
		fmt.Print(l.Val)
		fmt.Print(" ")
		l=l.Next
	}
	fmt.Println()
}

func main()  {

	head:=initNode(3)
	cur:=head
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(0)
	cur=cur.Next
	cur.Next=initNode(-4)
	printListNode(head)
	l:=hasCycle(head)
	fmt.Println(l)
}

func hasCycle(head *ListNode) bool {
	vis:=map[*ListNode]int{}
	for head!=nil{
		if _,ok:=vis[head];ok{
			return true
		}
		vis[head]=1
		head=head.Next
	}
	return false
}
