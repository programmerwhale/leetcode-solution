package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap2(height))
}

// 方法一：dp
// O(N),O(N)
func trap(height []int) int {
	leftmaxnums, rightmaxnums := []int{}, make([]int, len(height))
	maxx := 0
	for _, v := range height {
		if v > maxx {
			maxx = v
		}
		leftmaxnums = append(leftmaxnums, maxx)
	}
	maxx = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxx {
			maxx = height[i]
		}
		rightmaxnums[i] = maxx
	}
	res := 0
	for i := 0; i < len(height); i++ {
		res += min(leftmaxnums[i], rightmaxnums[i]) - height[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二：栈
func trap2(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		// 判断栈不为空 且当前高度高于栈顶元素的高度，才做处理
		// 两层循环
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 因为下面要再取一个值，所以这里判断一下
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			// 不知道左边高还是右边高
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		// 否则直接加入栈
		stack = append(stack, i)
	}
	return
}
