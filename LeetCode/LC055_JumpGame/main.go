package main

import "fmt"

/*
55. Jump Game跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

 

示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 

提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105
*/
func main()  {
	nums:=[]int{2,0,0}
	fmt.Println(jumpGame(nums))
}
// 时间复杂度O(N),空间复杂度O(1)
func jumpGame(nums []int)bool{
	// 只有一个元素，直接返回true
	if len(nums)==1{
		return true
	}
	// 当前所能覆盖到的最远下标
	cover:=0
	// i<len(nums)以防cover超出nums长度，nums[i]溢出
	for i:=0;i<=cover&&i<len(nums);i++{
		if i+nums[i]>cover{
			cover=i+nums[i]
		}
	}
	return cover>=len(nums)-1
}