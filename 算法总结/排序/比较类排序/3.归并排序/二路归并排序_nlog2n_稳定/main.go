package main

import "fmt"

//归并排序是一种稳定的排序方法。
// 和选择排序一样，归并排序的性能不受输入数据的影响，
// 但表现比选择排序好的多，因为始终都是O(nlogn）的时间复杂度。
// 代价是需要额外的内存空间。

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
	fmt.Println(sortArray(datas))
}
