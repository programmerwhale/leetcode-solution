package main

import (
	"fmt"
	"sort"
)

func main()  {
	ans:=[]int{1,3,2}
	fmt.Println(isMonotonic(ans))
}

// 暴力
// 两次遍历
func isMonotonic(nums []int) bool {
	if len(nums)==1{
		return true
	}
	return checkMonotonic(nums,1)||checkMonotonic(nums,-1)
}

func checkMonotonic(nums []int,val int) bool{
	for i:=1;i<len(nums);i++{
		if nums[i]!=nums[i-1]&&nums[i]-nums[i-1]<val{
			return false
		}
	}
	return true
}

// 优化 一次遍历
func isMonotonic2(nums []int) bool {
	if len(nums)==1{
		return true
	}
	inc, dec := true, true
	for i:=1;i<len(nums);i++{
		// 矛盾
		if nums[i-1]>nums[i]{
			inc=false
		}
		if nums[i-1]<nums[i]{
			dec=false
		}
	}
	return inc||dec
}

// 内置函数
func isMonotonic3(nums []int) bool {
	return sort.IntsAreSorted(nums) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}