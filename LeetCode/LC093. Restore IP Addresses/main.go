package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) (res []string) {
	var path []string
	backtracing(s,0,path,&res)
	return
}
func backtracing(s string,startindex int,path []string,res *[]string){
	if startindex==len(s)&&len(path)==4{
		tmp:=strings.Join(path,".")
		*res=append(*res,tmp)
	}
	for i:=startindex;i<len(s)&&i-startindex+1<=3;i++{
		if valid(s,startindex,i){
			path=append(path,s[startindex:i+1])
		}else{
			continue
		}
		backtracing(s,i+1,path,res)
		path=path[:len(path)-1]
	}
}
func valid(s string,startindex,end int)bool{
	if end!=startindex&&s[startindex]=='0'{
		//对于前导 0的情况
		return false
	}
	val,_:=strconv.Atoi(s[startindex:end+1])
	if val>=0&&val<256{
		return true
	}
	return false
}