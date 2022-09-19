package main

import "fmt"

func main()  {
	s:=[]string{"h","e","l","l","o"}
	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []string)  {
	n:=len(s)
	if n==0{
		return
	}
	for i:=0;i<n/2;i++{
		s[i],s[n-i-1]=s[n-i-1],s[i]
	}
}