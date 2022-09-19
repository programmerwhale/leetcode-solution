package main

import (
	"fmt"
)

func main()  {
	s := "  hello   world  "
	fmt.Println(reverseWords(s))
}
func reverseWords(s string) string {
	// 1.1先把开头和结尾的空格去掉
	start,end:=0,len(s)-1
	for i:=0;i<len(s);i++{
		if s[i]!=' '{
			start=i
			break
		}
	}
	for i:=len(s)-1;i>=0;i--{
		if s[i]!=' '{
			end=i
			break
		}
	}
	s=s[start:end+1]
	// 1.2再把中间的空格去掉
	b:=[]byte(s)
	index:=0
	for i:=0;i<len(b);i++{
		if b[i]!=' '||i>0&&b[i]==' '&&b[i-1]!=' '{
			b[index]=b[i]
			index++
		}
	}
	b=b[:index]
	n:=len(b)
	// 2.字符串反转
	for i:=0;i<n/2;i++{
		b[i],b[n-i-1]=b[n-i-1],b[i]
	}
	// 3.单词反转
	start,end=0,0
	for start<n&&end<n{
		for end<n&&b[end]!=' '{
			end++
		}
		// 防止出界
		sub:=b[start:min(len(b),end)]
		for i:=0;i<len(sub)/2;i++{
			sub[i],sub[len(sub)-i-1]=sub[len(sub)-i-1],sub[i]
		}
		// 此时end在空格上
		start=end+1
		end=start
	}
	return string(b)
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}