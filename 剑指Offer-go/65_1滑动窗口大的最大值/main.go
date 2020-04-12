package main

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 {
		return nums
	}

	max := nums[0]
	res := make([]int, 0)
	pos := 0
	for i := 1; i < k; i++ {
		max = getMax(max, nums[i])
		if max == nums[i] {
			pos = i
		}
	}
	res = append(res, nums[pos])

	left, right := 1, k

	for right <= len(nums)-1 {
		if pos == (left - 1) {
			pos = left
			for i := left + 1; i <= right; i++ {
				if nums[i] > nums[pos] {
					pos = i
				}
			}
		} else {
			if nums[pos] <= nums[right] {
				pos = right
			}
		}

		res = append(res, nums[pos])
		right++
		left++
	}

	return res
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {

}
