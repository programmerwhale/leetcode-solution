package main
/*
86. 分隔链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

 

示例 1：


输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]
示例 2：

输入：head = [2,1], x = 2
输出：[1,2]
 

提示：

链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/
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

	head:=initNode(1)
	cur:=head
	cur.Next=initNode(4)
	cur=cur.Next
	cur.Next=initNode(3)
	cur=cur.Next
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(5)
	cur=cur.Next
	cur.Next=initNode(2)
	printListNode(head)
	l:=partition(head,3)
	printListNode(l)
}


func partition(head *ListNode, x int) *ListNode {
	if head==nil{
		return head
	}
	l1,l2:=&ListNode{},&ListNode{}
	head1,head2:=l1,l2
	for head!=nil{
		if head.Val<x{
			l1.Next=initNode(head.Val)
			l1=l1.Next
		}else{
			l2.Next=initNode(head.Val)
			l2=l2.Next
		}
		head=head.Next
	}
	l1.Next=head2.Next
	return head1.Next
}