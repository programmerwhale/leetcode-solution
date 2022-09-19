package main
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

 

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
 

提示：
1 <= n <= 8

*/
func main()  {

}

// 回溯
// 时间空间复杂度O(n)
func generateParenthesis(n int) (res []string){
	path:=""
	backtracking(n,path,&res,0,0)
	return res
}
func backtracking(n int,path string,res *[]string,open,close int){
	if n*2==len(path){
		*res=append(*res,path)
		return
	}
	if open<n{
		path=path+"("
		backtracking(n,path,res,open+1,close)
		path=path[:len(path)-1]
	}
	// 注意右括号的结束条件，是要小于左括号数
	if close<open{
		path=path+")"
		backtracking(n,path,res,open,close+1)
		path=path[:len(path)-1]
	}
}