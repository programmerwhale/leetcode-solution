/*

 */
package main

import "fmt"

/*
84. 柱状图中最大的矩形
难度
困难

1770





给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。



示例 1:



输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：



输入： heights = [2,4]
输出： 4


提示：

1 <= heights.length <=105
0 <= heights[i] <= 104

*/
func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	maxArea := largestRectangleArea(heights)
	fmt.Println(maxArea)
}

// 暴力法-超时
func largestRectangleArea(heights []int) (maxArea int) {
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for j := i; j >= 0; j-- {
			if heights[j] >= heights[i] {
				left--
			} else {
				left = j
			}
		}
		for j := i; j < len(heights); j++ {
			if heights[j] >= heights[i] {
				right++
			} else {
				right = j
			}
		}
		area := heights[i] * (right - left + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 单调栈时空O(N)
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
