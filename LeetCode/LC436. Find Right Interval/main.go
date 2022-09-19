package main

import "sort"

func main()  {
	
}

// 方法一：
func findRightInterval(intervals [][]int) []int {
	//
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool { return intervals[i][0] >= p[1] })
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans
}

// 方法二：双指针
func findRightInterval2(intervals [][]int) []int {
	n := len(intervals)
	type pair struct{ x, i int }
	starts := make([]pair, n)
	ends := make([]pair, n)
	for i, p := range intervals {
		starts[i] = pair{p[0], i}
		ends[i] = pair{p[1], i}
	}
	sort.Slice(starts, func(i, j int) bool { return starts[i].x < starts[j].x })
	sort.Slice(ends, func(i, j int) bool { return ends[i].x < ends[j].x })

	ans := make([]int, n)
	j := 0
	for _, p := range ends {
		for j < n && starts[j].x < p.x {
			j++
		}
		if j < n {
			ans[p.i] = starts[j].i
		} else {
			ans[p.i] = -1
		}
	}
	return ans
}
