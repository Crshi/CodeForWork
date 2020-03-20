package main

import (
	"fmt"
)

//冒泡排序-升序
func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		flag := true
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				flag = false
				data[i], data[j] = data[j], data[i]
			}
		}

		if flag {
			return data
		}
	}

	return data
}

func main() {
	fmt.Println(BubbleSort([]int{10, 7, 8, 6, 3, 9}))
}
