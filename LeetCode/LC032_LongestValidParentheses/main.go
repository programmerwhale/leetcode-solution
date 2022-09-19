package main

import (
	"fmt"
)

/*
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。


示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：
输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'
*/
func main()  {
	s := "()(()"
	fmt.Println(longestValidParentheses2(s))
}

// 方法一：栈，加入一个特殊底元素
// 时间空间复杂度都是0(n)
func longestValidParentheses(s string) (maxx int) {
	if s==""{
		return
	}
	maxx=0
	q:=[]int{-1}
	// -1当个底
	for i,v:=range s{
		//进栈
		if v=='('{
			q=append(q,i)
		}else{
			//出栈
			// 如果此时括号都出栈了的话，-1被出了，现在的)需要顶替进入
			q=q[:len(q)-1]
			if len(q)==0 {
				q=append(q,i)
			}else{
				// 下标相减获得长度
				maxx=max(maxx,i - q[len(q)-1])
			}
		}
	}
	return
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

// 方法二：前后两遍遍历
// 时间复杂度0(N),空间复杂度0(1)
func longestValidParentheses2(s string) (maxx int) {
	if s==""{
		return
	}
	left,right:=0,0
	for _,v:=range s{
		if v=='('{
			left++
		}else{
			right++
		}
		if left==right{
			maxx=max(maxx,left+right)
		}else if right>left{//左到右，如果右边的大于左了肯定失效
			left,right=0,0
		}
	}
	left,right=0,0
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='('{
			left++
		}else{
			right++
		}
		if left==right{
			maxx=max(maxx,left+right)
		}else if left>right{//同理
			left,right=0,0
		}
	}
	return
}