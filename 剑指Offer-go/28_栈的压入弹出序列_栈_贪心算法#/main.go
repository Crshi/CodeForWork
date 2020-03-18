package main

import (
	"fmt"
)

func validateStackSequences(pushed []int, popped []int) bool {
	datas := make([]int, 0)
	length := len(pushed)
	j := 0
	for i := 0; i < length; i++ {
		datas = append(datas, pushed[i])
		for len(datas) != 0 && j < length && datas[len(datas)-1] == popped[j] {
			datas = datas[:len(datas)-1]
			j++
		}
	}

	return len(datas) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
