package main

import (
	"fmt"
	"math"
	"sort"
)

/*
面试题 01.01. 判定字符是否唯一
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

*/
func main() { fmt.Println(maximumCandies([]int{5, 7, 8, 6}, 8)) }
func maximumCandies(candies []int, k int64) int {
	n := int64(len(candies))
	sum := int64(0)
	minn := math.MaxInt64
	for _, v := range candies {
		sum += int64(v)
		if v < minn {
			minn = v
		}
	}
	if sum < k {
		return 0
	}
	if n >= k {
		sort.Ints(candies)
		return candies[n-k]
	}
	tmp := int(sum / k)
	childrenCnt := int64(0)
	for i := tmp; i > 0; i-- {
		for _, v := range candies {
			childrenCnt += int64(v) / int64(i)
			if childrenCnt >= k {
				return i
			}
		}
		childrenCnt = int64(0)
	}
	return -1

}
