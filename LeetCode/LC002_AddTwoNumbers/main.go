package main

import "fmt"

/*
2. 两数相加
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
 

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

*/

func main()  {
	// 构造链表
	l1:=initNode(2)
	cur:=l1
	cur.Next=initNode(4)
	cur=cur.Next
	cur.Next=initNode(3)

	l2:=initNode(5)
	cur=l2
	cur.Next=initNode(6)
	cur=cur.Next
	cur.Next=initNode(4)

	printListNode(l1)
	printListNode(l2)
	// 调用 输出
	l3:=addTwoNumbers(l1,l2)
	printListNode(l3)
}

// 链表结构初始化
type ListNode struct {
	Val int
	Next *ListNode
}

// 初始化节点
func initNode(val int) *ListNode{
	return &ListNode{val,nil}
}

// 时间复杂度O(max(m,n)),两个链表中的最大长度
// 空间复杂度O(1)，返回值不计入空间复杂度。
func addTwoNumbers(l1 *ListNode,l2 *ListNode)*ListNode{
	// 记录进位
	carry:=0
	head:=&ListNode{}
	cur:=head
	for l1!=nil||l2!=nil||carry!=0{
		next:=&ListNode{}
		if l1!=nil{
			carry+=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			carry+=l2.Val
			l2=l2.Next
		}
		next.Val=carry%10
		carry/=10
		cur.Next=next
		cur=cur.Next
	}
	return head.Next
}

// 链表输出
func printListNode(l *ListNode){
	for l!=nil{
		fmt.Print(l.Val)
		fmt.Print(" ")
		l=l.Next
	}
	fmt.Println()
}