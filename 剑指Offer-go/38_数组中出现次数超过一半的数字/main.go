package main

func majorityElement(nums []int) int {
	value := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != value {
			count--
			if count == 0 {
				value = nums[i]
				count = 1
			}
		} else {
			count++
		}
	}

	return value
}

func main() {
	majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2})
}
