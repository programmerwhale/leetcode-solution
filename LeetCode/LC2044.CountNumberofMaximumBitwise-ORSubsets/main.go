package main

func main()  {
	
}

// todo
// 方法一：位运算
// 时间复杂度O(2^N*N)，空间复杂度O(1)
func countMaxOrSubsets(nums []int) (ans int) {
	maxOr:=0
	for i:=1;i<1<<len(nums);i++{
		or:=0
		for j,num:=range nums{
			if i>>j&1==1{
				or|=num
			}
		}
		if or>maxOr{
			maxOr=or
			ans=1
		}else if or==maxOr {
			ans++
		}
	}
	return
}

// 方法二：回溯
// 时间复杂度O(2^N)，空间复杂度O(N)
func countMaxOrSubsets2(nums []int) (ans int) {
	maxOr:=0
	var dfs func(int,int)
	dfs=func(pos,or int){
		if pos==len(nums){
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0,0)
	return
}

// dp