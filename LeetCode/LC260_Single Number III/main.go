package main

import "fmt"

func main() {
	ans := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(ans))
}

func singleNumber(nums []int) (ans []int) {
	xorsum := 0
	for _, num := range nums {
		xorsum ^= num
	}
	// !!!!!
	// x&(-x) 找到最右边1的位置
	index := xorsum & -xorsum
	x1, x2 := 0, 0
	for _, num := range nums {
		// 分为0 1两组，每组各一个出现一次的数
		if num&index == 0 {
			x1 ^= num
		} else {
			x2 ^= num
		}
	}
	return []int{x1, x2}
}
