/*
9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
示例 4：

输入：x = -101
输出：false
 

提示：

-231 <= x <= 231 - 1
 

进阶：你能不将整数转为字符串来解决这个问题吗？

 */
package main

import (
	"fmt"
	"strconv"
)

func main()  {
	res:=isPalindrome2(12)
	fmt.Println(res)
}

/*
转字符串：
1. 负数可以直接返回false
2. int转字符串，两两对比
时间复杂度：O(n)
空间复杂度：O(1)?
*/
func isPalindrome1(x int)bool{
	if x<0{
		return  false
	}
	str:=strconv.Itoa(x)
	n:=len(str)
	for i:=0;i<len(str);i++{
		if str[i]!=str[n-i-1]{
			return false
		}
	}
	return true
}

/*
不转字符串的实现：
1.思路：把后半部分的数字取出来反转，前后两个数字对比是否相等；
2.结束条件：当后>=前的时候可以结束取值反转, 即后<新时一直循环；
3.奇数偶数的问题，可以在最后返回时处理；如果数字位数为奇数，反转数字后，后半段数字/10就忽略了中间这个多出来的数；
4.以0结尾和负数的直接可以返回false；
*/
func isPalindrome2(x int)bool{
	if x<0||(x!=0&&x%10==0){
		return  false
	}
	reverNum:=0
	for reverNum<x{
		reverNum=reverNum*10+x%10
		x/=10
	}
	return reverNum==x||reverNum/10==x
}
