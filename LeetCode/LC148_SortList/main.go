package main
/*
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

 

示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]
 

提示：

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
 

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
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
	head:=initNode(4)
	cur:=head
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(1)
	cur=cur.Next
	cur.Next=initNode(3)
	printListNode(head)
	l:=sortList(head)
	printListNode(l)
}

// 分治法，左右分别排序再合并
// 时间复杂度：O(nlogn)，其中 nn 是链表的长度。
//
//空间复杂度：O(logn)，其中 nn 是链表的长度。空间复杂度主要取决于递归调用的栈空间
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func findMid(head, tail *ListNode)*ListNode{
	slow,fast:=head,head
	// 注意这里结束处理
	for fast!=tail{
		slow=slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return slow
}

func mergeList(list1,list2  *ListNode)*ListNode {
	var dummy *ListNode
	dummy=initNode(0)
	cur:=dummy
	l1,l2:=list1,list2
	for l1!=nil&&l2!=nil{
		if l1.Val<l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else{
			cur.Next=l2
			l2=l2.Next
		}
		cur=cur.Next
	}
	if l1!=nil{
		cur.Next=l1
	}
	if l2!=nil{
		cur.Next=l2
	}
	return dummy.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	mid := findMid(head,tail)
	return mergeList(sort(head, mid), sort(mid, tail))
}
