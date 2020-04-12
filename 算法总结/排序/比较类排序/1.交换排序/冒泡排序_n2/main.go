package main

import (
	"fmt"
)

//冒泡排序-升序
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		fmt.Println(nums)

		if flag == false {
			break
		}
	}
	return nums
}

func main() {
	datas := []int{10, 1, 35, 61, 89, 36, 55}
	// start := time.Now()
	BubbleSort(datas)
	fmt.Println(datas)
	// fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
}
