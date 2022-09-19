package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums:=[]int{1,1,2}
	fmt.Println(permuteUnique(nums))
}
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var used []bool
	used=make([]bool,len(nums))
	sort.Ints(nums)
	backtracking(&res,&used,[]int{},nums)
	return res
}
func backtracking(res *[][]int,used *[]bool,path,nums []int){
	if len(path)==len(nums){
		temp:=make([]int,len(path))
		copy(temp,path)
		*res=append(*res,temp)
	}
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]&&(*used)[i-1]==false||(*used)[i]==true{
			continue
		}
		(*used)[i]=true
		path=append(path,nums[i])
		backtracking(res,used,path,nums)
		path=path[:len(path)-1]
		(*used)[i]=false
	}
}