package main

import "fmt"

func main()  {
	nums := []int{5,3,4,4,7,3,6,11,8,5,11}
	fmt.Println(totalSteps(nums))
}

func totalSteps(nums []int) (res int) {
	type pair struct{ v, t int }
	stack := []pair{}
	for i:=0;i<len(nums);i++{
		cnt:=0
		for len(stack)>0&&nums[i]>=stack[len(stack)-1].v{
			cnt = max(cnt, stack[len(stack)-1].t)
			stack=stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			cnt++
			res = max(res, cnt)
		} else {
			cnt = 0
		}
		stack=append(stack,pair{nums[i],cnt})
	}
	return
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
