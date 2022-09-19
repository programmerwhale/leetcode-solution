package main

import "fmt"

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
func numUniqueEmails(emails []string) int {
	mp := map[string]int{}
	for _, email := range emails {
		tmp := make([]byte, 0)
		for i := 0; i < len(email); i++ {
			if email[i] == '@' {
				tmp = append(tmp, email[i:]...)
			}
			if email[i] == '.' {
				continue
			}
			if email[i] == '+' {
				for email[i+1] != '@' {
					i++
				}
				continue
			}
			fmt.Println(string(tmp))
			tmp = append(tmp, email[i])
		}
		mp[string(tmp)]++
	}
	return len(mp)
}
