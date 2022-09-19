package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println(fourSum(nums, 8))
}
func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		//对n1的去重
		// 用if continue
		if i > 0 && nums[i-1] == n1 {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				//对n2的去重，用for
				for l < r && n2 == nums[l] {
					l++
				}
				//对n3的去重，用for
				for l < r && n3 == nums[r] {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return
}

func fourSum(nums []int, target int) (res [][]int) {
	n := len(nums)
	if n < 4 {
		return
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n1 := nums[i]
		for j := i + 1; j < n; j++ {
			if j > 0 && nums[j] == nums[j-1] {
				continue
			}
			n2 := nums[j]
			l, r := j+1, n-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 == target {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for r > l && nums[r] == n4 {
						r--
					}
				} else if n1+n2+n3+n4 < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return
}
