package main

func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	return n
}

func main() {

}
