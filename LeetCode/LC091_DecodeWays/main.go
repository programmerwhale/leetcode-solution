package main

import (
	"fmt"
	"strconv"
)

func main()  {
	s:="2101"
	fmt.Println(numDecodings(s))
}

// dp
func numDecodings(s string) int {
	if len(s)==1 && s[0]=='0'{
		return 0
	}
	dp:=make([]int,len(s)+1)
	dp[0]=1
	for i:=1;i<=len(s);i++{
		if s[i-1]!='0'{
			dp[i]+=dp[i-1]
		}
		if i>1&&isValid(s[i-2:i]){
			dp[i]+=dp[i-2]
		}
	}
	return dp[len(s)]
}


// 优化dp
func numDecodings2(s string) int {
	if len(s)==1 && s[0]=='0'{
		return 0
	}
	// 多申请一个空间
	n1,n2,n3:=0,1,0
	for i:=1;i<=len(s);i++{
		// n3在这需要置为0
		n3=0
		if s[i-1]!='0'{
			n3+=n2
		}
		if i>1&&isValid(s[i-2:i]){
			n3+=n1
		}
		n1,n2=n2,n3
	}
	return n3
}

func isValid(s string)bool{
	n:=len(s)
	if n==2 && s[0]=='0'{
		return false
	}
	v,_:=strconv.Atoi(s)
	if v>=1&&v<=26{
		return true
	}
	return false
}