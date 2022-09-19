package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(longestPalindrome("ba"))
}
// 1.暴力（超时）
func longestPalindrome(s string) (res string) {
	n:=len(s)
	if n==1{
		return s
	}
	res=string(s[:1])
	maxx:=math.MinInt64
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			length:=j-i+1
			if check(s,i,j)&&length>maxx{
				maxx=length
				res=string(s[i:j+1])
			}
		}
	}
	return res
}

func check(s string,i,j int)bool{
	for i<j{
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}

//2.中心扩展
// 时间复杂度O(N^2)
// 空间复杂度O(1)
func longestPalindrome2(s string) string {
	n:=len(s)
	if n==1{
		return s
	}
	start,end:=0,0
	for i:=0;i<n;i++{
		//bab
		left1,right1:=expandAroundCenter(s,i,i)
		//baab
		left2,right2:=expandAroundCenter(s,i,i+1)
		if right1-left1>end-start{
			start,end=left1,right1
		}
		if right2-left2>end-start{
			start,end=left2,right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	// 左右不过界 && 左右相等
	for ;left>=0&&right<len(s)&&s[left]==s[right];left,right=left-1,right+1{}
	// 记得还原
	return left+1,right-1
}
// todo
// 3.Manacher算法
func longestPalindrome3(s string) string {
	return s
}