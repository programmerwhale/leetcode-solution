package main

func main()  {
	
}

type ListNode struct {
    Val int
    Next *ListNode
}

// 方法一：转成数组判断是否为回文数组
// 时间空间复杂度O(N)
func isPalindrome(head *ListNode) bool {
    ans:=[]int{}
    for head!=nil{
        ans=append(ans,head.Val)
        head=head.Next
    }
    i,j:=0,len(ans)-1
    for i<j {
        if ans[i]!=ans[j]{
            return false
        }
        i++
        j--
    }
    return true
}

// 方法二：快慢指针
// 整个流程可以分为以下五个步骤：
//
//找到前半部分链表的尾节点。
//反转后半部分链表。
//判断是否回文。
//恢复链表。
//返回结果。
//
// 时间复杂度O(N),空间复杂度O(1)
func isPalindrome2(head *ListNode) bool {
    if head==nil||head.Next==nil{
        return true
    }
    mid:=findMidNode(head)
    halfstart:=reverseListNode(mid.Next)
    p1,p2:=head,halfstart
    for p1!=nil&&p2!=nil{
        if p1.Val!=p2.Val{
            return false
        }
        p1,p2=p1.Next,p2.Next
    }
    return true
}

func findMidNode(head *ListNode)*ListNode{
    fast,slow:=head,head
    // 如果不能蹦两下就不蹦了
    for fast.Next!=nil&&fast.Next.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }
    return slow
}

func reverseListNode(head *ListNode)*ListNode{
    var pre *ListNode
    cur:=head
    for cur!=nil{
        next:=cur.Next
        cur.Next=pre
        pre=cur
        cur=next
    }
    return pre
}