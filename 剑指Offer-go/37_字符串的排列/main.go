package main

import (
	"fmt"
)

func permutation(s string) []string {
	datas := []byte(s)
	res := make([]string, 0)
	maps := make(map[string]bool, 0)
	recr(maps, datas, "")
	for index, _ := range maps {
		res = append(res, index)
	}
	return res
}

func recr(maps map[string]bool, s []byte, tmp string) {
	if len(s) == 0 {
		maps[tmp] = true
		return
	}

	for i := 0; i < len(s); i++ {
		if i == 0 {
			// t := make([]byte, len(s)-1)
			// copy(t, s[1:])
			recr(maps, s[1:], tmp+string(s[i]))
			continue
		}

		if i == len(s)-1 {
			// t := make([]byte, len(s)-1)
			// copy(t, s[:len(s)-1])
			recr(maps, s[:len(s)-1], tmp+string(s[i]))
			continue
		}

		t1 := make([]byte, i)
		copy(t1, s[:i])

		recr(maps, append(t1, s[i+1:]...), tmp+string(s[i]))
	}
}

func main() {
	fmt.Println(permutation("aab"))
}
