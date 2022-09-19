package main

func main()  {
	
}
// 时间复杂度O(N^2)
// 空间复杂度O(N)
func combine(n int, k int) (res [][]int) {
	var path []int
	backtrace(&res,&path,1,n,k)
	return
}
func backtrace(res *[][]int,path *[]int,start int,n int, k int){
	if len(*path)==k{
		tmp:=make([]int,k)
		copy(tmp,*path)
		*res=append(*res,tmp)
		return
	}
	// 剪枝
	for i:=start;i<=n-(k-len(*path))+1;i++{
		*path=append(*path,i)
		// 下一个start是i+1
		backtrace(res,path,i+1,n,k)
		*path=(*path)[:len(*path)-1]
	}
}