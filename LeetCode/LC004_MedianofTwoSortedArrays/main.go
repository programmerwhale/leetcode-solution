package main

import "fmt"

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/
func main(){
	ans1:=[]int{1,2,3,4,5}
	ans2:=[]int{1,2,3}
	fmt.Println(findMedianSortedArrays(ans1,ans2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m,n:=len(nums1),len(nums2)
	i,j:=0,0
	index:=0
	k:=(m+n)>>1
	temp:=0
	mid,mid2:=0,0
	for (i<m-1||j<n-1)&&index<=k{
		if nums1[i]<=nums2[j]{
			i++
			temp=nums1[i]
		}else{
			j++
			temp=nums1[j]
		}
		if index==k-1{
			mid=temp
		}
		if index==k{
			mid2=temp
		}
		index++
	}
	if k%2==1{
		return float64(mid)
	}
	return float64((mid+mid2)>>1)
}