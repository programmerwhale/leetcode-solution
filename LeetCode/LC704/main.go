/*
题目描述：
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

提示：
你可以假设 nums中的所有元素是不重复的。
n将在[1, 10000]之间。
nums的每个元素都将在[-9999, 9999]之间。

思路：


*/
package main

import "fmt"

func main()  {
	ans:=[]int{-1,0,3,5,9,12}
	target:=search(ans,9)
	fmt.Println(target)
}

func search(ans []int, target int)int{
	l,r:=0,len(ans)
	for l<=r{
		mid:=l+(r-l)/2
		if ans[mid]==target{
			return mid
		}else if ans[mid]<target{
			l=mid+1
		}else{
			r=mid-1
		}
	}
	return -1
}
