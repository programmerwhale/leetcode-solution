package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "1"
	num2 := "9"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) (res string) {
	str1 := []byte(num1) //long
	str2 := []byte(num2)
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	carry := 0
	i, n := 0, len(str2)
	for carry > 0 || i < n {
		n1, n2 := 0, 0
		// 注意这两句界限
		if i < len(str1) {
			n1, _ = strconv.Atoi(string(str1[len(str1)-i-1]))
		}
		if i < n {
			n2, _ = strconv.Atoi(string(str2[n-i-1]))
		}
		tmp := n1 + n2 + carry
		carry = tmp / 10
		tmpstr := strconv.Itoa(tmp % 10)
		res = tmpstr + res
		i++
	}
	if i < len(str1) {
		res = string(str1[:len(str1)-i]) + res
	}
	return res
}

func addStrings2(num1 string, num2 string) (res string) {
	carry:=0
	for len(num1)>0||len(num2)>0||carry>0{
		temp:=0
		if len(num1)>0{
			v:=num1[len(num1)-1]-'0'
			temp+=int(v)
			num1=num1[:len(num1)-1]
		}
		if len(num2)>0{
			v:=num2[len(num2)-1]-'0'
			temp+=int(v)
			num2=num2[:len(num2)-1]
		}
		temp+=carry
		res=strconv.Itoa(temp%10)+res
		carry=temp/10
	}
	return
}
