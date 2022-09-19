package main

import "fmt"

func main()  {
	s:="We are happy."
	fmt.Println(replaceSpace(s))
}
func replaceSpace(s string) string {
	space:=0
	for _,v:=range s{
		if v==' '{
			space++
		}
	}
	tmp:=make([]byte,space*2)
	b:=[]byte(s)
	length:=len(b)
	b=append(b,tmp...)
	i:=len(b)-1
	index:=length-1
	for i>=0&&index>=0{
		if b[index]!=' '{
			b[i]=b[index]
			i--
		}else{
			b[i]='0'
			b[i-1]='2'
			b[i-2]='%'
			i-=3
		}
		index--
	}
	return string(b)
}

// 另一种写法
func replaceSpace2(s string) string {
	cnt := 0
	for _, v := range s {
		if v == ' ' {
			cnt++
		}
	}
	// 不是乘3
	substr := make([]byte, cnt*2)
	s = string(substr) + s

	b := []byte(s)
	index := 0
	// 注意起始位置
	for i := cnt * 2; i < len(b); i++ {
		if b[i] != ' ' {
			b[index] = b[i]
			index++
		} else {
			b[index] = '%'
			b[index+1] = '2'
			b[index+2] = '0'
			index += 3
		}
	}
}