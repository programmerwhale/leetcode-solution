package main

import "fmt"

func main()  {
	fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring(s string) int {
	// 初始值赋-1，
	// 需要考虑全程都没有重复字符的情况
	temp,res:=-1,0
	mp:=map[byte]int{}
	str:=[]byte(s)
	for i,v:=range str{
		if p,ok:=mp[v];ok{
			temp=max(temp,p)
		}
		mp[v]=i
		res=max(res,i-temp)
	}
	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
