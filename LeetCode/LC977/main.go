/*
977. 有序数组的平方
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

 

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
 

提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序

思路：
因为有时间复杂度的要求，所以只能遍历一次，考虑使用双指针
循环条件的左右指针可以相等，因为会有单数的情况，如果不能相等的话，就会把中间的漏掉
*/
package main

import "fmt"

func main()  {
	nums:=[]int{-7,-3,2,3,11}
	res:=sortedSquares(nums)
	fmt.Println(res)
}

func sortedSquares(nums []int)(ans []int){
	n:=len(nums)
	i,j:=0,n-1
	index:=n-1
	ans=make([]int,n)
	for i<=j{
		a,b:=nums[i]*nums[i],nums[j]*nums[j]
		if a>b{
			ans[index]=a
			i++
		}else{
			ans[index]=b
			j--
		}
		index--
	}
	return ans
}
