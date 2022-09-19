package main

import "fmt"

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	index:=len(s)-1
	for index-k>=0{
		str:=""
		str=s[index-k+1:index+1]
		/*
		if index+1==k{
			str=s[index-k+1:]
		}else{
			str=s[index-k+1:index]
		}*/
		if hashSub(str,power,modulo,hashValue)==true{
			return str
		}else{
			index--
		}
	}
	return ""
}

func hashSub(s string, power int, modulo int,hashValue int)bool{

	sum:=0
	powers:=1
	for _,v:=range s{
		sum+= powers * int(v-96)%modulo
		powers*=power%modulo
		sum%=modulo
	}
	return sum%modulo==hashValue
}

func main(){
	fmt.Println(subStrHash("xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk",22,51,41,9))
}
