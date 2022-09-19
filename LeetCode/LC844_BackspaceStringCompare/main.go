package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

// 方法一：
// 时间空间复杂度o(m+n)
func backspaceCompare(s string, t string) bool {
	return realstr(s) == realstr(t)
}
func realstr(s string) string {
	b := []byte{}
	for _, ch := range s {
		if ch != '#' {
			b = append(b, byte(ch))
		} else if len(b) > 0 {
			b = b[:len(b)-1]
		}
	}
	return string(b)
}

// 方法二：双指针
// 时间复杂度o(m+n)
// 空间复杂度o(1)
func backspaceCompare2(s string, t string) bool {
	skips, skipt := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skips++
				i--
			} else if skips > 0 {
				skips--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipt++
				j--
			} else if skipt > 0 {
				skipt--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
