package main

func createTargetArray(nums []int, index []int) []int {
	res := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if index[i] == 0 {
			res = append([]int{nums[i]}, res[:]...)
		} else {
			if index[i] == len(res) {
				t := nums[i]
				res = append(res, t)
			} else {
				rear := append([]int{nums[i]}, res[index[i]:]...)
				res = append(res[:index[i]], rear...)
			}
		}

	}

	return res
}

func main() {
	t := createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})

	println(t)
}
