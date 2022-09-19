package main

import (
	"fmt"
	"strings"
)

func main()  {
	s:="YazaAay"
	fmt.Println(longestNiceSubstring(s))
}

//  思路1：双指针
func longestNiceSubstring(s string) (ans int) {
	if s==""{
		return 0
	}
	s=strings.ToLower(s)
	i,j:=0,0
	maxlen:=0
	//maxStart:=0
	for i<=j&&j<len(s)-1{
		if s[i]==s[j]{
			j++
		}else{
			if j-i>maxlen{
				maxlen=j-i
			}
			i=j
		}
	}
	return maxlen
}
