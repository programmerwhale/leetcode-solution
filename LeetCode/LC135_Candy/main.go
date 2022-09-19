package main

import (
	"fmt"
)

/*
135. Candy分发糖果
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

 

示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 

提示：

n == ratings.length
1 <= n <= 2 * 104
0 <= ratings[i] <= 2 * 104
*/
func main()  {
	ratings:=[]int{1,2,87,87,87,2,1}
	fmt.Println(candy(ratings))
}

// 两次遍历
// 第一次从前往后：只要左<右，右糖果为左糖果数+1
// 第二次从后往前：只要左>右，左为max(左当前值，右现有糖果数+1)
// 时间复杂度0(N),空间复杂度O(n)
func candy(rationgs []int)(res int){
	n:=len(rationgs)
	candys:=make([]int,n)
	// 第一次从前往后：只要左<右，右糖果为左糖果数+1
	for i:=0;i<n-1;i++{
		if rationgs[i]<rationgs[i+1]{
			candys[i+1]=candys[i]+1
		}
	}
	// 第二次从后往前：只要左>右，左为max(左当前值，右现有糖果数+1)
	for i:=n-1;i>0;i--{
		if rationgs[i-1]>rationgs[i]{
			// 左为max(左当前值，右现有糖果数+1)
			candys[i-1]=max(candys[i]+1,candys[i-1])
		}
	}
	fmt.Println(candys)
	for _,v:=range candys{
		res+=v
	}
	// 最后加个总数，就不用遍历设置初始值了
	return res+n
}

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}

// 一次遍历，不占用额外空间
// 时间复杂度0(N),空间复杂度O(1)
func candy2(ratings []int) int {
	n := len(ratings)
	// 第一个孩子已就位，加入递增序列
	ans, inc, dec, pre := 1, 1, 0, 1
	// 从第二个孩子开始
	for i := 1; i < n; i++ {
		// 增，等于的也算增
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			// 减
			dec++
			// 如果等长了，递增的最后元素要划到递减里
			if dec == inc {
				dec++
			}
			// 递减的长度整体抬1
			ans += dec
			pre = 1
		}
	}
	return ans
}