package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}
func calPoints(ops []string) int {
	stack := []int{}
	for _, ch := range ops {
		val, error := strconv.Atoi(ch)
		if error == nil {
			stack = append(stack, val)
		}
		switch ch {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			tmp := stack[len(stack)-1]
			stack = append(stack, tmp*2)
		case "+":
			if len(stack) > 1 {
				tmp := stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, tmp)
			}
		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}
