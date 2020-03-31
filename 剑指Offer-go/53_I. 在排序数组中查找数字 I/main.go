package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := (left + right) / 2
	count := 0
	for left <= right {
		if nums[middle] == target {
			count++
			for i := middle + 1; i < len(nums); i++ {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}

			for i := middle - 1; i >= 0; i-- {
				if nums[i] == target {
					count++
				} else {
					break
				}
			}

			return count
		}

		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
		middle = (left + right) / 2
	}

	return count
}

func main() {
	fmt.Println(search([]int{0, 8, 8, 8, 8, 8}, 8))
}
