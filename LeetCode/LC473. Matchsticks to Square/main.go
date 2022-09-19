package main

import "sort"

func main()  {
	nums:=[]int{3,2,2,1}
	makesquare(nums)
}
func makesquare(matchsticks []int) bool {
	n:=len(matchsticks)
	sort.Ints(matchsticks)
	for i:=0;i<n/2;i++{
		matchsticks[i],matchsticks[n-1-i]=matchsticks[n-1-i],matchsticks[i]
	}
	sum:=0
	for _,v:=range matchsticks{
		sum+=v
	}
	if sum%4!=0{
		return false
	}
	side:=sum/4
	edges := [4]int{}

	var dfs func(int)bool
	dfs=func(index int)bool{
		if index==n{
			return true
		}
		for i,_:=range edges{
			// 每次新元素都从第一条边开始试
			edges[i]+=matchsticks[index]
			// 只有可以继续往下放的才会继续递归，可以等于，不然就没法往下走了
			// 不能直接在if里加matchsticks[index]，因为如果可以放进去的就不减掉了
			if edges[i]<=side&&dfs(index+1){
				return true
			}
			// 回溯
			edges[i]-=matchsticks[index]
		}
		// 一旦有放不进去的就会返回false
		return false
	}
	return dfs(0)
}