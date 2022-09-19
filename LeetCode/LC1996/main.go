/*
你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] = [attacki, defensei] 表示游戏中第 i 个角色的属性。

如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。更正式地，如果认为角色 i 弱于 存在的另一个角色 j ，那么 attackj > attacki 且 defensej > defensei 。

返回 弱角色 的数量。

示例 1：

输入：properties = [[5,5],[6,3],[3,6]]
输出：0
解释：不存在攻击和防御都严格高于其他角色的角色。
示例 2：

输入：properties = [[2,2],[3,3]]
输出：1
解释：第一个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
示例 3：

输入：properties = [[1,5],[10,4],[4,3]]
输出：1
解释：第三个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。

思路：
1. 根据第一个值排序，再对比第二个值
 左右两个值排序规则向反，左边从大到小，右边从小到大，这样只记住最大值就可以了，最大值出现之后 后面的都是需要的数
 注意去记小于的
时间复杂度O(nlogn),空间复杂度O(logn)，其中n为数组的长度。排序时使用的栈空间为 O(logn)。
2.
*/

package main

import (
	"fmt"
	"sort"
)

func main()  {
	properties:=[][]int{}
	properties=append(properties, []int{9,4})
	properties=append(properties, []int{8,5})
	properties=append(properties, []int{10,1})
	fmt.Println(properties)
	fmt.Println("---------------")
	//sort2(properties)
	ans:=numberOfWeakCharacters(properties)
	fmt.Println("---------------")
	fmt.Println(ans)
}

func numberOfWeakCharacters(properties [][]int) (ans int) {
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] > q[0] || p[0] == q[0] && p[1] < q[1]
	})
	//sort2(properties)
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return
}
func PrintProperties(properties [][]int){
	for i:=0;i<len(properties);i++{
		fmt.Println(properties[i])
	}
}
func sort2(p [][]int){
	for i:=0;i<=len(p);i++{
		for j:=len(p)-1;j>=1;j--{
			if p[j][0]>p[j-1][0]||p[j][0]==p[j-1][0]&&p[j][1]<p[j-1][1]{
				p[j],p[j-1]=p[j-1],p[j]
			}
		}
	}
	fmt.Println(p)
}
