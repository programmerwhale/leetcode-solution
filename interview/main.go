package main

import "strconv"

func main() {
}
func calculate(s string) int {
	if len(s) == 0 {
		return -1
	}
	b := []byte(s)
	stack := []int{}
	var num int
	flag := 1
	flag2, flag3 := false, false
	for i := 0; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			num, _ = strconv.Atoi(string(b[i]))
			num *= flag
			flag = 1
			if flag2 == true {
				//乘法，取值计算再压入栈
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num *= tmp
				flag2 = false
			}
			if flag3 == true {
				//除法，取值计算再压入栈
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				num = tmp / num
				flag2 = false
			}
			stack = append(stack, num)
		}
		if b[i] == '-' {
			flag = -1
		}
		if b[i] == '*' {
			flag2 = true
		}
		if b[i] == '/' {
			flag3 = true
		}
	}
	res := 0
	for len(stack) > 0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}
