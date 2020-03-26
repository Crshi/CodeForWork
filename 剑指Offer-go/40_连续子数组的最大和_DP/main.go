package main

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		res = max(nums[i], res)
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
