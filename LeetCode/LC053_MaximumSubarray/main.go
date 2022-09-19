/*
53. Maximum Subarray
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

 

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23
 

提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
 

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/
package main

import "fmt"

func main()  {
	ans:=[]int{5,4,-1,7,8}
	fmt.Println(maxSubArray2(ans))
}
// 方法一:滑动窗口
// 时间复杂度O（n），空间复杂度O(1)
func maxSubArray(nums []int) (res int) {
	res=nums[0]
	if len(nums)==1{
		return
	}
	for i:=1;i<len(nums);i++{
		if nums[i-1]+nums[i]>nums[i]{
			nums[i]+=nums[i-1]
		}
		if res<nums[i]{
			res=nums[i]
		}
	}
	return
}

// 方法二：分治-线段树
// 时间复杂度O（nlogn），空间复杂度O(logn)
/*
	将问题分解为最大子序和都在左边/右边/跨中间
*/
func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

type Status struct {
	/*
		lSum 表示 [l,r] 内以 l 为左端点的最大子段和
		rSum 表示 [l,r] 内以 r 为右端点的最大子段和
		mSum 表示 [l,r] 内的最大子段和
		iSum 表示 [l,r] 的区间和
	*/
	lSum, rSum, mSum, iSum int
}

// 合并
func pushUp(l, r Status) Status {
	// 合并后左边的最大子段和=max(左侧最大子段和，左侧区间全部的和+右侧以 l 为左端点的最大子段和)
	lSum := max(l.lSum, l.iSum + r.lSum)
	// 合并后右边的最大子段和=max(右侧最大子段和，右侧区间全部的和+左侧以 r 为右端点的最大子段和)
	rSum := max(r.rSum, r.iSum + l.rSum)
	// 合并跨中间的最大子段和=max(max(左右测区间的最大子段和),左侧以 r 为右端点的最大子段和+右侧以 l 为左端点的最大子段和)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	// 区间和：左+右
	iSum := l.iSum + r.iSum
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if (l == r) {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	// 分治，将跨中间的分解为左边和右边，中间的跟着左边
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}