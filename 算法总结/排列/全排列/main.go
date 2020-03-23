package main

import (
	"fmt"
)

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	dict := map[string]bool{}
	str := []byte(s)
	var f func(index int)
	f = func(index int) {
		if index == len(str) {
			dict[string(str)] = true
			return
		}
		for i := index; i < len(str); i++ {
			tmp := str[index]
			str[index] = str[i]
			str[i] = tmp
			f(index + 1)
			str[i] = str[index]
			str[index] = tmp
		}
	}
	f(0)
	res := []string{}
	for k, _ := range dict {
		res = append(res, k)
	}
	return res
}

func main() {
	t := permutation("abc")
	fmt.Println(t)
}
