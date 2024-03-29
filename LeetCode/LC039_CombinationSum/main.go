package main

import (
	"fmt"
	"sort"
)

/*
39. 组合总和Combination Sum

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 

对于给定的输入，保证和为 target 的不同组合数少于 150 个。

 

示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []
 

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500
*/
func main()  {
	candidates:=[]int{2,3,6,7}
	res:=combinationSum(candidates,7)
	fmt.Println(res)
}
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sort.Ints(candidates)
	used:=make(map[int]bool)
	backtracking(&res,path,candidates,target,0,0,used)
	return res
}

// 不允许重复
/*
加了一个map或者数组uesd 记录在该枝有没有用到，key为candidates的下标
在回溯前需要排序，如果当前值和前一个值相同，并且上一个值已经取值了，就可以取
举例（1,1）；
只取第一个1和只取第二个1，虽然不是同一个1，但是path会重复，（1，x，x），（1，x，x）；
所以如果要取第二个1，就是两个1都要取，前一个1没取值，第二个1不可以取，（1，1，x）；
*/
// 时间复杂度：O(S)，其中 S 为所有可行解的长度之和。O(n×2^n)是一个比较松的上界,但实际运行情况是远远小于这个上界的。
// 空间复杂度：O(n)。除答案数组外，我们需要 O(n) 的空间存储used。
// 优化方法：剪枝--sum + candidates[i] <= target
func backtracking(res *[][]int,path []int,candidates []int,target int,startindex,sum int,used map[int]bool){
	// 结束条件
	if sum==target{
		temp:=make([]int,len(path))
		copy(temp,path)
		*res=append(*res,temp)
	}
	if sum>target{
		return
	}
	// 优化后
	for i:=startindex;i<len(candidates)&&sum + candidates[i] <= target;i++{
	// 优化前
	//for i:=startindex;i<len(candidates);i++{
		// 检查是否和上一个是否重复
		if i>0&&candidates[i]==candidates[i-1]&&used[i-1]==false{
			continue
		}
		used[i]=true
		sum+=candidates[i]
		path=append(path,candidates[i])
		// startIndex变为i+1，从下一元素开始
		backtracking(res,path,candidates,target,i+1,sum,used)
		sum-=candidates[i]
		path=path[:len(path)-1]
		used[i]=false
	}
	return
}

// 允许重复
// 时间复杂度：O(S)，其中 S 为所有可行解的长度之和。O(n×2^n)是一个比较松的上界,但实际运行情况是远远小于这个上界的。
// 空间复杂度：O(target)。除答案数组外，空间复杂度取决于递归的栈深度，在最差情况下需要递归 O(target) 层。
// 优化方法：剪枝--sum + candidates[i] <= target
func backtracking2(res *[][]int,path []int,candidates []int,target int,startindex,sum int){
	// 结束条件
	if sum==target{
		temp:=make([]int,len(path))
		copy(temp,path)
		*res=append(*res,temp)
	}
	if sum>target{
		return
	}
	// 优化后？？好像有点问题
	for i:=startindex;i<len(candidates)&&sum + candidates[i] <= target;i++{
		// 优化前
		//for i:=startindex;i<len(candidates);i++{
		sum+=candidates[i]
		path=append(path,candidates[i])
		backtracking2(res,path,candidates,target,i,sum)
		sum-=candidates[i]
		path=path[:len(path)-1]
	}
	return
}