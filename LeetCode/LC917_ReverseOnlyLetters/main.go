package main

import (
	"fmt"
	"unicode"
)

/*
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。

 

示例 1：

输入：s = "ab-cd"
输出："dc-ba"
示例 2：

输入：s = "a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入：s = "Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"
 

提示

1 <= s.length <= 100
s 仅由 ASCII 值在范围 [33, 122] 的字符组成
s 不含 '\"' 或 '\\'

*/

func main()  {
	fmt.Println(reverseOnlyLetters("ab-cd"))
}
// 时间复杂度O(N),空间复杂度O(1)
// 双指针，注意循环条件
// 判断是否为字母：unicode.IsLetter(rune(s[left]))
func reverseOnlyLetters(s string) string {
	ans:=[]byte(s)
	n:=len(s)
	left,right:=0,n-1
	for left<right{
		for left<right&&!unicode.IsLetter(rune(s[left])) {
			left++
		}
		for left<right&&!unicode.IsLetter(rune(s[right])) {
			right--
		}
		if left>=right{
			break
		}
		ans[left],ans[right]=ans[right],ans[left]
		left++
		right--
	}
	return string(ans)
}