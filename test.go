package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	datas := make(map[byte]int, 0)
	left := 0
	right := 0
	res := 0
	for left <= right && right < len(s) {
		if datas[s[right]] == 0 {
			datas[s[right]]++
			right++
			res = max(res, right-left)
		} else {
			datas[s[left]]--
			left++
		}
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
