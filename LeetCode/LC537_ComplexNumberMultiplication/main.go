package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	num1 := "1+1i"
	num2 := "1+1i"
	fmt.Println(complexNumberMultiply(num1,num2))
}

func complexNumberMultiply(num1 string, num2 string) string  {
	a1:=strings.Split(num1,"+")
	a2:=strings.Split(num2,"+")
	reala,_:=strconv.Atoi(a1[0])
	realb,_:=strconv.Atoi(a2[0])
	realasum:=reala*realb
	imaginaryia,_:=strconv.Atoi(a1[1][:len(a1[1])-1])
	imaginaryib,_:=strconv.Atoi(a2[1][:len(a2[1])-1])
	imaginaryisum:=imaginaryia*imaginaryib*(-1)
	sum:=realasum+imaginaryisum
	temp:=reala*imaginaryib+realb*imaginaryia
	return strconv.Itoa(sum)+"+"+strconv.Itoa(temp)+"i"
}
