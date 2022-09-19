package main

import "fmt"

func main(){
	nums:=[]int{1,2,1}
	fmt.Println(nextGreaterElements(nums))
}
func nextGreaterElements(nums []int) (res []int) {
	stack:=[]int{}
	res=make([]int,len(nums))
	n:=len(nums)
	for i:=n*2-1;i>=0;i--{
		for len(stack)>0{
			if nums[i%n]>=stack[len(stack)-1]{
				stack=stack[:len(stack)-1]
			}else{
				break
			}
		}
		if len(stack)>0{
			res[i%n]=stack[len(stack)-1]
		}else{
			res[i%n]=-1
		}
		stack=append(stack,nums[i%n])
	}
	return
}