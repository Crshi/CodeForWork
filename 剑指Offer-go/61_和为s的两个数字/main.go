package main

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left != right {
		tmp := nums[left] + nums[right]
		if tmp > target {
			right--
		} else {
			if tmp < target {
				left++
			} else {
				return []int{nums[left], nums[right]}
			}
		}
	}

	return []int{}
}

func main() {

}
