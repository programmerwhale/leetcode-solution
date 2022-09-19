package main

import (
	"fmt"
	"math"
)

/*
134. Gas Station加油站
在一条环路上有N个加油站，其中第i个加油站有汽油gas[i]升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。
你从其中的一个加油站出发，开始时油箱为空。

如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

说明: 

如果题目有解，该答案即为唯一答案。
输入数组均为非空数组，且长度相同。
输入数组中的元素均为非负数。
示例 1:

输入:
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

输出: 3

解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。
示例 2:

输入:
gas  = [2,3,4]
cost = [3,4,3]

输出: -1

解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。
*/
func main()  {
	gas:=[]int{3,1,1}
	cost:=[]int{1,2,2}
	fmt.Println(canCompleteCircuit3(gas,cost))
}
// 暴力法（超时）
// 时间复杂度：O(n^2)，空间复杂度：O(n)
func canCompleteCircuit(gas []int, cost []int) int {
	for i:=0;i<len(cost);i++{
		// 初始剩余油量
		rest:=gas[i]-cost[i]
		// index为下一个加油站的位置；
		// 如果从最后一个加油站开始的话；循环，回到开头；
		index:=(i+1)%len(cost)
		// 以i开头跑一圈
		for rest>0 && index!=i{
			rest+=gas[index]-cost[index]
			index=(index+1)%len(cost)
		}
		// 如果还有余量，并且回到起始点时；与上面循环结束条件相对应
		if rest>=0&&index==i{
			return i
		}
	}
	return -1
}

// 暴力法优化：
// 考虑以下情况
// 1. 如果cost总和小于gas总和，一定跑不完一圈
// 2. 先从第一个起点跑，如果余量出现负数，一定跑不完一圈；如果没负数，返回起点0
// 3. 2出现负数后，从后向前，看哪个节点能这个负数填平，能把这个负数填平的节点就是出发节点。
// 时间复杂度：O(n) ,空间复杂度：O(1)
func canCompleteCircuit2(gas []int, cost []int) int {
	curSum:=0
	minn:=math.MaxInt64
	for i:=0;i<len(gas);i++{
		rest:=gas[i]-cost[i]
		curSum+=rest
		if curSum<minn{
			minn=curSum
		}
	}
	// 情况1
	if curSum<0{
		return -1
	}

	// 情况2
	if minn>=0{
		return 0
	}

	// 情况3
	for i:=len(gas)-1;i>0;i-- {
		rest:=gas[i]-cost[i]
		// 看能不能填平
		minn+=rest
		if minn>=0{
			return i
		}
	}
	return -1
}

// 贪心算法
// 起始加油站的剩余油量一定要大于0，跑的时候如果发现小于0，更换起始点
// 时间复杂度：O(n) ,空间复杂度：O(1)
func canCompleteCircuit3(gas []int, cost []int) int{
	total:=0
	start:=0
	rest:=0
	for i:=0;i<len(cost);i++{
		// 累计剩余油量
		rest+=gas[i]-cost[i]
		total+=gas[i]-cost[i]
		// 一旦剩余油量为负数，抛弃掉
		if rest<0&&i<len(gas)-1{
			start=i+1
			// 剩余油量清零
			rest=0
		}
	}
	// 如果一圈总油量小于cost，一定跑不完一圈
	if total<0{
		return -1
	}
	return start
}