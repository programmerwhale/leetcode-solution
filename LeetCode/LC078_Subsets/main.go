package main

import "fmt"

/*
78. 子集Subsets
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

 

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]
 

提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/
func main()  {
	nums:=[]int{1,2,3}
	fmt.Println(Subsets(nums))
}

func Subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	backtracking(&res,path,nums,0)
	return res
}

func backtracking(res *[][]int,path []int,nums []int,startindex int){
	// 注意要申请长度
	temp:=make([]int,len(path))
	copy(temp,path)
	*res=append(*res,temp)
	for i:=startindex;i<len(nums);i++{
		path=append(path,nums[i])
		// startindex为当前下标+1
		backtracking(res,path,nums,i+1)
		path=path[:len(path)-1]
	}
	return
}
