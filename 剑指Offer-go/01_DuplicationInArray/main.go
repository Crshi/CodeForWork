package main

import "fmt"

func main() {
	var nums = [7]int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums[:]))

}

func findRepeatNumber(nums []int) int {
	length := len(nums)
	for _,x := range nums {

		if x >= length {
			x = x - length
		}

		if nums[x] >= length {
			return x
		}

		nums[x] = nums[x] + length
	}

	return -1
}
