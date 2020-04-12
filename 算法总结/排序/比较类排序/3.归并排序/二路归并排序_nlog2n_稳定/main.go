package main

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	left := sortArray(nums[0:middle])
	right := sortArray(nums[middle:])

	return merge(left, right)
}

func merge(l, r []int) []int {
	i, j := 0, 0
	m, n := len(l), len(r)
	var res []int
	for i < m && j < n {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, r[j:]...)

	return res
}

func main() {
	datas := []int{10, 1, 35, 61, 89, 36, 55}
	sortArray(datas)
	fmt.Println(datas)
}
