/*
209. 长度最小的子数组
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105

进阶：

如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
---------------------------------------------------------------------
思路：
1. 双指针滑动窗口
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main()  {
	ans:=[]int{2,3,1,2,4,3}
	a:=minSubArrayLen2(7,ans)
	fmt.Println(a)
}

// 滑动窗口，时间复杂度O(N),空间复杂度O(1)
func minSubArrayLen(target int, nums []int) int {
	n:=len(nums)
	if n==1{
		if nums[0]==target {
			return 1
		}
		return 0
	}
	sum,cnt:=nums[0],n
	i,j:=0,0
	for j<n{
		if sum<=target{
			sum+=nums[j]
			if sum==target{
				cnt=min(cnt,j-i+1)
			}
			j++
		}else if sum>target{
			sum-=nums[i]
			i++
		}
	}
	return cnt
}

// 前缀和，时间复杂度O(Nlogn),空间复杂度O(1)
func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n + 1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i - 1] + nums[i - 1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound - (i - 1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}

