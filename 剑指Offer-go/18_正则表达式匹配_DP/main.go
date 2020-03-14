package main

import "fmt"

func isMatch(s string, p string) bool {
	//双指针

	i,j := 0,0

	for i < len(s) {

		if j >= len(p) {
			return false
		}

		if p[j] == '.' {
			i ++
			j ++
			continue
		}

		if p[j] == '*' {
			//p[j-1]可以出现任意次
			for {
				if i < len(s) && s[i] == p[j-1] {
					i++
					continue
				} else {
					break
				}
			}
			j++
			continue
		}

		if s[i] == p[j] {
			i++
			j++
			continue
		} else {
			if j+1 < len(p) && p[j+1] =='*' {
				j++
				continue
			}
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isMatch("mississippi","mis*is*p*."))
}
