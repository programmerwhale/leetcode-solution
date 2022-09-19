package main
/*
21. Merge Two Sorted Lists合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/
import "fmt"

func main()  {
	list1:=initNode(1)
	cur:=list1
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(4)
	printListNode(list1)

	list2:=initNode(1)
	cur=list2
	cur.Next=initNode(3)
	cur=cur.Next
	cur.Next=initNode(4)
	printListNode(list2)

	l3:=mergeTwoLists(list1,list2)
	printListNode(l3)
}

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head:=&ListNode{}
	cur:=head
	for list1!=nil&&list2!=nil{
		// 记得申请一个next
		next:=&ListNode{}
		if list1.Val<list2.Val{
			next.Val=list1.Val
			list1=list1.Next
		}else{
			next.Val=list2.Val
			list2=list2.Next
		}
		cur.Next=next
		cur=cur.Next
	}
	// 在最后处理多出来的部分，因为在循环里要取值，不能为空
	if list1!=nil{
		cur.Next=list1
	}
	if list2!=nil{
		cur.Next=list2
	}
	return head.Next
}