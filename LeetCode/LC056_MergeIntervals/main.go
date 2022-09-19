package main

import (
	"fmt"
	"sort"
)

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

 

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/
func main()  {
	intervals:=[][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(intervals))
}
// 排序，然后先加入结果集再编辑
// 时间复杂度0(nlogn)，空间0(logn)
func merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	res=append(res,intervals[0])
	for i:=1;i<len(intervals);i++{
		if intervals[i][0]<=res[len(res)-1][1]{
			res[len(res)-1][1]=max(res[len(res)-1][1],intervals[i][1])
		}else{
			res=append(res,intervals[i])
		}
	}
	return res
}
func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}