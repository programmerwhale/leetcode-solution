package main

import (
	"fmt"
	"math"
)

/*
120. 三角形最小路径和
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。



示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：
输入：triangle = [[-10]]
输出：-10

提示：
1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

进阶：
你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
*/

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal3(triangle))
}

// 二维dp
// 时间复杂度O(n^2)，空间复杂度O(n^2)
func minimumTotal(triangle [][]int) int {
	m, n := len(triangle), len(triangle[len(triangle)-1])
	if m == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	res := math.MaxInt64
	for _, v := range dp[n-1] {
		if v < res {
			res = v
		}
	}
	return res
}

// 原数组上改
// 时间复杂度O(n^2)，空间复杂度O(1)
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	for i := 1; i < n; i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
	res := math.MaxInt64
	for _, v := range triangle[n-1] {
		if v < res {
			res = v
		}
	}
	return res
}

// 一维 需要倒着来，因为左边的值已经被修改，会导致错误答案
// 时间复杂度O(n^2)，空间复杂度O(n)
func minimumTotal3(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}
	dp := make([]int, m)
	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			// 注意这里下标是j
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	res := math.MaxInt64
	for _, v := range dp {
		if v < res {
			res = v
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
