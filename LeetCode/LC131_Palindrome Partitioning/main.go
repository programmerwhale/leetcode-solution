package main

func main()  {
	
}

// 记忆化搜索+回溯
func partition(s string) (res [][]string) {
	var path []string
	// 备忘录：0 代表未计算 1代表是回文 2 代表不是回文
	memo:=make([][]int,len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}
	backtracing(s,0,&path,&res,memo)
	return
}

func backtracing(s string,startindex int,path *[]string,res *[][]string,memo [][]int){
	if startindex==len(s){
		tmp:=make([]string,len(*path))
		copy(tmp,*path)
		*res=append(*res,tmp)
		return
	}
	for i:=startindex;i<len(s);i++{
		if memo[startindex][i] == 2 { //不回文，直接跳过
			continue
		}
		//记录为回文 或没有记录但isPali调用为真
		if memo[startindex][i]==1||isPalindrome(s,startindex,i,memo){
			*path=append(*path,s[startindex:i+1])
		}else{
			continue
		}
		backtracing(s,i+1,path,res,memo)
		*path=(*path)[:len(*path)-1]
	}
}

func isPalindrome(s string,start,end int,memo [][]int)bool{
	i,j:=start,end
	for ;i<j;i,j=i+1,j-1{
		if s[i]!=s[j]{
			return false
		}
	}
	memo[i][j] = 1
	return true
}